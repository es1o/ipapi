# ipapi

Super simple api that returns some info about IP address.

## sample response

```bash
{
  "city": {
    "city":"Gdansk",
    "country_code":"PL",
    "latitude":11.11,
    "longitude":11.11,
    "postcode":"XX-XXX",
    "state1":"Pomerania",
    "state2":"",
    "timezone":"Europe/Warsaw"
  },
  "ip": "A.B.C.D",
  "record": {
    "anonymous_ip": {
      "is_anonymous": false,
      "is_anonymous_vpn": false,
      "is_hosting_provider": false,
      "is_public_proxy": false,
      "is_tor_exit_node": false
    },
    "autonomous_system_number": 5617,
    "autonomous_system_organization": "TPNET Orange Polska Spolka Akcyjna",
    "continent": {
      "code": "EU",
      "geoname_id": 6255148,
      "names": {
        "en": "Europe"
      }
    },
    "country": {
      "geoname_id": 798544,
      "iso_code": "PL",
      "names": {
        "en": "Poland"
      }
    },
    "has_anonymous_ip": false,
    "registered_country": {
      "geoname_id": 798544,
      "iso_code": "PL",
      "names": {
        "en": "Poland"
      }
    },
    "rir": "ripencc"
  }
}
```


## How to run

1. Download database from ip66.dev
2. Mount databae as /ip66.mmdb inside docker container
3. Optional: download city database, e.g. GeoLite2-City and mount in /city.mmdb
4. App listens on port 8080