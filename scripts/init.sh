#!/bin/sh

HOST="elasticsearch1"
PORT='9200'

# Create index: iller
curl --no-progress-meter -XPUT "http://$HOST:$PORT/iller" -H 'Content-Type: application/json' -d '{
  "mappings": {
    "_doc": {
      "properties": {
        "il": { "type": "text" },
        "nufus": { "type": "integer" },
        "ilceler": { "type": "nested" }
      }
    }
  }
}'

# Add ten sample documents to "iller" index
curl -XPOST "http://$HOST:$PORT/iller/_bulk" -H 'Content-Type: application/x-ndjson' -d '
{ "index" : { "_index" : "iller", "_id" : "1" } }
{ "il" : "istanbul", "nufus" : 15655924, "ilceler" : ["avcilar", "besiktas", "esenler"] }
{ "index" : { "_index" : "iller", "_id" : "2" } }
{ "il" : "ankara", "nufus" : 5445026, "ilceler" : ["cankaya", "etimesgut", "mamak"] }
{ "index" : { "_index" : "iller", "_id" : "3" } }
{ "il" : "izmir", "nufus" : 4279477, "ilceler" : ["dikili", "menemen", "odemis"] }
{ "index" : { "_index" : "iller", "_id" : "4" } }
{ "il" : "bursa", "nufus" : 2936803, "ilceler" : ["nilufer", "osmangazi", "yildirim"] }
{ "index" : { "_index" : "iller", "_id" : "5" } }
{ "il" : "adana", "nufus" : 2201670, "ilceler" : ["cukurova", "saricam", "seyhan"] }
{ "index" : { "_index" : "iller", "_id" : "6" } }
{ "il" : "sanliurfa", "nufus" : 2213964, "ilceler" : ["akcakale", "bozova", "harran"] }
{ "index" : { "_index" : "iller", "_id" : "7" } }
{ "il" : "konya", "nufus" : 2130544, "ilceler" : ["akoren", "cihanbeyli", "eregli"] }
{ "index" : { "_index" : "iller", "_id" : "8" } }
{ "il" : "antalya", "nufus" : 321457, "ilceler" : ["alanya", "kepez", "konyaalti"] }
{ "index" : { "_index" : "iller", "_id" : "9" } }
{ "il" : "aydin", "nufus" : 1161702, "ilceler" : ["efeler", "soke", "kusadasi"] }
{ "index" : { "_index" : "iller", "_id" : "10" } }
{ "il" : "tekirdag", "nufus" : 1167059, "ilceler" : ["cerkezkoy", "corlu", "ergene"] }
'

# Create index: ulkeler
curl --no-progress-meter -XPUT "http://$HOST:$PORT/ulkeler" -H 'Content-Type: application/json' -d '{
  "mappings": {
    "_doc": {
      "properties": {
        "ulke": { "type": "text" },
        "nufus": { "type": "integer" },
        "baskent": { "type": "text" }
      }
    }
  }
}'

# Add ten sample documents to "ulkeler" index
curl -XPOST "http://$HOST:$PORT/ulkeler/_bulk" -H 'Content-Type: application/x-ndjson' -d '
{ "index" : { "_index" : "ulkeler", "_id" : "1" } }
{ "ulke" : "turkiye", "nufus" : 85372377, "baskent" : "ankara" }
{ "index" : { "_index" : "ulkeler", "_id" : "2" } }
{ "ulke" : "amerika_birlesik_devletleri", "nufus" : 336080688, "baskent" : "washington" }
{ "index" : { "_index" : "ulkeler", "_id" : "3" } }
{ "ulke" : "cin", "nufus" : 1409670000, "baskent" : "pekin" }
{ "index" : { "_index" : "ulkeler", "_id" : "4" } }
{ "ulke" : "hindistan", "nufus" : 1412604531, "baskent" : "yeni_delhi" }
{ "index" : { "_index" : "ulkeler", "_id" : "5" } }
{ "ulke" : "rusya", "nufus" : 146424729, "baskent" : "moskova" }
{ "index" : { "_index" : "ulkeler", "_id" : "6" } }
{ "ulke" : "brezilya", "nufus" : 203062512, "baskent" : "brasilia" }
{ "index" : { "_index" : "ulkeler", "_id" : "7" } }
{ "ulke" : "guney_kore", "nufus" : 51439038, "baskent" : "seul" }
{ "index" : { "_index" : "ulkeler", "_id" : "8" } }
{ "ulke" : "iran", "nufus" : 85231028, "baskent" : "tahran" }
{ "index" : { "_index" : "ulkeler", "_id" : "9" } }
{ "ulke" : "birlesik_krallik", "nufus" : 67026292, "baskent" : "londra" }
{ "index" : { "_index" : "ulkeler", "_id" : "10" } }
{ "ulke" : "birlesik_arap_emirlikleri", "nufus" : 9282410, "baskent" : "abu_dabi" }
'