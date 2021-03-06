package crud

import (
	"context"
	"fmt"
	"sync"

	"github.com/graph-gophers/dataloader"

	"github.com/spaceuptech/space-cloud/gateway/model"
	"github.com/spaceuptech/space-cloud/gateway/utils"
)

type resultsHolder struct {
	sync.Mutex
	results []*dataloader.Result
	metas   []meta
}

type meta struct {
	whereClause map[string]interface{}
	op          string
	dbType      string
}

type queryResult struct {
	doc      interface{}
	metaData *model.SQLMetaData
}

func (holder *resultsHolder) getResults() []*dataloader.Result {
	holder.Lock()
	defer holder.Unlock()

	return holder.results
}

func (holder *resultsHolder) addResult(i int, result *dataloader.Result) {
	holder.Lock()
	holder.results[i] = result
	holder.Unlock()
}

func (holder *resultsHolder) getWhereClauses() []interface{} {
	holder.Lock()
	defer holder.Unlock()

	arr := make([]interface{}, 0)
	for _, v := range holder.metas {
		arr = append(arr, v.whereClause)
	}
	return arr
}

func (holder *resultsHolder) addMeta(op, dbType string, whereClause map[string]interface{}, matchClause []map[string]interface{}) {
	holder.Lock()
	for i, where := range matchClause {
		for k, v := range where {
			if k == "$or" {
				k = fmt.Sprintf("%s:%d", k, i)
			}
			whereClause[k] = v
		}
	}
	holder.metas = append(holder.metas, meta{whereClause: whereClause, op: op, dbType: dbType})
	holder.Unlock()
}

func (holder *resultsHolder) fillResults(metData *model.SQLMetaData, res []interface{}) {
	holder.Lock()
	defer holder.Unlock()

	// Create a where clause index
	index := 0

	length := len(holder.results)
	for i := 0; i < length; i++ {

		// Continue if result already has a value
		if holder.results[i] != nil {
			continue
		}

		// Get the where clause
		meta := holder.metas[index]
		isOperationTypeOne := meta.op == utils.One
		docs := make([]interface{}, 0)
		for _, doc := range res {
			if utils.Validate(meta.dbType, meta.whereClause, doc) {
				docs = append(docs, doc)
			}
			if isOperationTypeOne {
				break
			}
		}

		// Increment the where clause index
		index++

		var result interface{}
		if isOperationTypeOne {
			if len(docs) > 0 {
				result = docs[0]
			} else {
				result = nil
			}
		} else {
			result = docs
		}
		// Store the matched docs in result
		holder.results[i] = &dataloader.Result{Data: queryResult{doc: result, metaData: metData}}
	}
}

func (holder *resultsHolder) fillErrorMessage(err error) {
	holder.Lock()

	length := len(holder.results)
	for i := 0; i < length; i++ {
		if holder.results[i] == nil {
			holder.results[i] = &dataloader.Result{Error: err}
		}
	}
	holder.Unlock()
}

func (m *Module) getLoader(key string) (*dataloader.Loader, bool) {
	m.dataLoader.dataLoaderLock.RLock()
	defer m.dataLoader.dataLoaderLock.RUnlock()
	loader, ok := m.dataLoader.loaderMap[key]
	return loader, ok
}

func (m *Module) createLoader(key string) *dataloader.Loader {
	m.dataLoader.dataLoaderLock.Lock()
	defer m.dataLoader.dataLoaderLock.Unlock()
	// DataLoaderBatchFn is the batch function of the data loader
	cache := &dataloader.NoCache{}
	loader := dataloader.NewBatchedLoader(m.dataLoaderBatchFn, dataloader.WithCache(cache))
	m.dataLoader.loaderMap[key] = loader
	return loader
}

func (m *Module) dataLoaderBatchFn(c context.Context, keys dataloader.Keys) []*dataloader.Result {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	var dbAlias, col string

	// Return if there are no keys
	if len(keys) == 0 {
		return []*dataloader.Result{}
	}

	holder := resultsHolder{
		results: make([]*dataloader.Result, len(keys)),
		metas:   make([]meta, 0),
	}

	for index, key := range keys {
		req := key.(model.ReadRequestKey)

		dbAlias = req.DBAlias
		col = req.Col

		// Execute query immediately if it has options
		if req.HasOptions {
			// Add task to wait group
			wg.Add(1)

			go func(i int) {
				defer wg.Done()

				// make sures metric get collected for following read request
				req.Req.IsBatch = false      // NOTE: DO NOT REMOVE THIS
				req.Req.Options.Select = nil // Need to make this nil so that we load all the fields data
				// Execute the query
				res, metaData, err := m.Read(ctx, dbAlias, req.Col, &req.Req, req.ReqParams)
				if err != nil {

					// Cancel the context and add the error response to the result
					cancel()
					holder.addResult(i, &dataloader.Result{Error: err})
					return
				}

				// Add the response to the result
				holder.addResult(i, &dataloader.Result{Data: queryResult{doc: res, metaData: metaData}})
			}(index)

			// Continue to the next key
			continue
		}

		// Append the where clause to the list
		holder.addMeta(req.Req.Operation, req.DBType, req.Req.Find, req.Req.MatchWhere)
	}

	// Wait for all results to be done
	wg.Wait()

	clauses := holder.getWhereClauses()

	// Fire the query only if where clauses exist
	if len(clauses) > 0 {
		// Prepare a merged request
		req := model.ReadRequest{Find: map[string]interface{}{"$or": clauses}, Operation: utils.All, Options: &model.ReadOptions{}}
		// Fire the merged request
		res, metaData, err := m.Read(ctx, dbAlias, col, &req, model.RequestParams{Resource: "db-read", Op: "access", Attributes: map[string]string{"project": m.project, "db": dbAlias, "col": col}})
		if err != nil {
			holder.fillErrorMessage(err)
		} else {
			holder.fillResults(metaData, res.([]interface{}))
		}
	}

	// do some async work to get data for specified keys
	// append to this list resolved values
	return holder.getResults()
}
