# HA

```
curl \
                                        -H "Content-Type: application/json" \
                                        -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI1ODkzNDI1MjQzNzM0ZDMxYjQyMDIyN2E3ZDhlOWNiNyIsImlhdCI6MTY5ODA4MzExNCwiZXhwIjoyMDEzNDQzMTE0fQ.5EbEvsAPPswD6184CwuIYTgdY63iYHmXC5j8WZJrHos" \
                                        -d '{
                                   "entity_id": "media_player.yandex_station_742078e2880c08040810",
                                   "media_content_id": "Лиза",
                                   "media_content_type": "text"
                                  }' \
                                        http://192.168.0.99:8123/api/services/media_player/play_media
```

```
curl \
                                        -H "Content-Type: application/json" \
                                        -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI1ODkzNDI1MjQzNzM0ZDMxYjQyMDIyN2E3ZDhlOWNiNyIsImlhdCI6MTY5ODA4MzExNCwiZXhwIjoyMDEzNDQzMTE0fQ.5EbEvsAPPswD6184CwuIYTgdY63iYHmXC5j8WZJrHos" \
                                        -d '{
                                   "entity_id": "media_player.yandex_station_742078e2880c08040810",
                                   "media_content_id": "Включи свет",
                                   "media_content_type": "command"
                                  }' \
                                        http://192.168.0.99:8123/api/services/media_player/play_media
```