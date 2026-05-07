package handlers

// GET /analytics/:code
//
// Response:
//   {
//     "code": "aB3xYz",
//     "total_clicks": 1234,
//     "by_day": [...],
//     "by_country": [...]
//   }
//
// Reads from Postgres aggregated tables populated by the worker.

// Analytics is the gin.HandlerFunc for GET /analytics/:code.
// TODO: implement using gin.Context once dependencies are wired.
func Analytics() {}
