# TODO 

## 1. Log Format Management
* [x] Create parser with fields
* [x] Get one parser by name
* [x] Get all/list parsers

## 2. Parser Engine Core
* [x] Implement regex parsing logic (named groups extraction)
* [x] Implement JSON parsing logic (key mapping)
* [x] Type conversion (string, int, float, datetime)
* [x] Handle required field validation
* [ ] Error handling for malformed logs

## 3. PostgreSQL Log Storage
* [ ] Create logs table schema 
* [ ] Insert parsed logs batch operation
* [ ] Indexing for performance (timestamp, parser_id)

## 4. Log Upload Pipeline
* [ ] File upload endpoint (multipart/form-data)
* [ ] Parse uploaded file line by line
* [ ] Batch insert to PostgreSQL
* [ ] Return upload stats (processed/failed lines)

## 5. Query API
* [ ] Basic filtering endpoint (time range, field filters)
* [ ] Pagination support (limit/offset)
* [ ] Sort by timestamp/fields
* [ ] Count/stats endpoint

## 6. Docker Setup
* [ ] Complete docker-compose.yml (PostgreSQL + App only)
* [ ] Dockerfile for Go app
* [ ] Environment variables configuration

## 7. Demo & Documentation
* [ ] Sample log files (Apache access + JSON app logs)
* [ ] Seed default parsers
* [ ] README with curl examples
* [ ] Postman collection

---

## 🔐 1. **Log Format Management**
Resource: `log-formats`

### ➕ Create (Already Done)
```http
POST /log-formats
```
**Body:**
```json
{
  "name": "apache-access",
  "is_json": false,
  "regex_pattern": "(?P<ip>\\S+) - - \\[(?P<timestamp>[^\\]]+)\\] \"(?P<method>\\S+) (?P<path>\\S+) (?P<protocol>\\S+)\" (?P<status>\\d+) (?P<size>\\d+)",
  "fields": [
    {
      "raw_name": "ip",
      "semantic_name": "client_ip", 
      "type": "string",
      "required": true
    },
    {
      "raw_name": "timestamp",
      "semantic_name": "request_time",
      "type": "datetime",
      "datetime_format": "02/Jan/2006:15:04:05 -0700",
      "required": true
    },
    {
      "raw_name": "method",
      "semantic_name": "http_method",
      "type": "string",
      "required": true
    },
    {
      "raw_name": "status",
      "semantic_name": "status_code",
      "type": "int",
      "required": true
    }
  ]
}
```

### 📄 Get One
```http
GET /log-formats/{name}
```

### 📃 List All  
```http
GET /log-formats
```

---

## 📥 2. **Log Upload Pipeline**
Resource: `logs`

### 📤 Upload Log File
```http
POST /logs/upload
```
**Multipart/form-data:**
* `file`: the log file (.log, .txt, .json)
* `parser_id`: UUID of parser to use

```bash
curl -F "file=@access.log" -F "parser_id=abc-123" http://localhost:8080/logs/upload
```

**Returns:**
```json
{
  "lines_processed": 1000,
  "lines_failed": 5,
  "duration_ms": 800,
  "parser_used": "apache-access"
}
```

---

## 🔎 3. **Query API**  
Resource: `logs`

### 🧾 Query Logs
```http
POST /logs/query
```
**Body:**
```json
{
  "parser_id": "abc-123",
  "filters": {
    "timestamp_from": "2025-01-01T00:00:00Z",
    "timestamp_to": "2025-01-02T00:00:00Z", 
    "status_code": "500",
    "client_ip": "192.168.1.1"
  },
  "limit": 50,
  "offset": 0,
  "sort_by": "request_time",
  "sort_order": "desc"
}
```

**Returns:**
```json
{
  "results": [
    {
      "request_time": "2025-01-01T12:00:00Z",
      "client_ip": "192.168.1.1", 
      "http_method": "GET",
      "status_code": 500,
      "raw_log": "192.168.1.1 - - [01/Jan/2025:12:00:00 +0000] \"GET /api/users HTTP/1.1\" 500 1234"
    }
  ],
  "total": 25,
  "page_info": {
    "limit": 50,
    "offset": 0,
    "has_more": false
  }
}
```

### 📊 Analytics/Stats
```http
GET /logs/stats?parser_id=abc-123&hours=24
```
**Returns:**
```json
{
  "total_logs": 10000,
  "error_rate": 2.5,
  "top_status_codes": {
    "200": 8500,
    "404": 1000, 
    "500": 500
  },
  "logs_per_hour": [
    {"hour": "2025-01-01T12:00:00Z", "count": 500},
    {"hour": "2025-01-01T13:00:00Z", "count": 750}
  ]
}
```

---

## Sample Files for Testing

### Apache Access Log (`sample_access.log`)
```
192.168.1.1 - - [01/Jan/2025:12:00:00 +0000] "GET /index.html HTTP/1.1" 200 1024
192.168.1.2 - - [01/Jan/2025:12:01:00 +0000] "POST /api/login HTTP/1.1" 500 256  
```

### JSON App Log (`sample_app.json`)
```json
{"timestamp": "2025-01-01T12:00:00Z", "level": "ERROR", "message": "Database connection failed", "user_id": 123}
{"timestamp": "2025-01-01T12:01:00Z", "level": "INFO", "message": "User logged in", "user_id": 456}
```
