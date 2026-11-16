
# FR05 — Notes

## Part A — Step 2: HTTP request and response

| Question | Answer |
|---|---|
| Request method | GET |
| Path | `/v1/forecast` |
| Query parameter names | `latitude`, `longitude`, `current` |
| Status code and text | 200 OK |
| Content-Type | application/json |

## Part A — Step 3: JSON description

### JSON response

```json
{
  "current": {
    "temperature_2m": 8.4,
    "time": "2026-11-16T14:30",
    "wind_speed_10m": 12.6
  },
  "current_units": {
    "temperature_2m": "°C",
    "time": "iso8601",
    "wind_speed_10m": "km/h"
  },
  "latitude": 48.89,
  "longitude": 2.2,
  "timezone": "Europe/Paris"
}
```

### JSON types and Go types

| JSON value | JSON type | Go type |
|---|---|---|
| `latitude` | number | `float64` |
| `longitude` | number | `float64` |
| `timezone` | string | `string` |
| `current` | object | `struct` |
| `current.temperature_2m` | number | `float64` |
| `current.time` | string | `string` |
| `current.wind_speed_10m` | number | `float64` |
| `current_units` | object | `struct` |
| `current_units.temperature_2m` | string | `string` |
| `current_units.time` | string | `string` |
| `current_units.wind_speed_10m` | string | `string` |

## Part A — Step 4: API error tests

| Test | Request | Expected result | Observed result |
|---|---|---|---|
| Missing parameters | `/v1/forecast` | 400 Bad Request | 400 Bad Request |
| Invalid latitude | `/v1/forecast?latitude=abc&longitude=2.2` | 400 Bad Request | 400 Bad Request |
| Unknown endpoint | `/v1/previsions` | 404 Not Found | 404 Not Found |
