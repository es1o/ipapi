# ipapi

Super simple API that returns some info about an IP address.

## Sample response

```json
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

## Configuration

`API_KEY` — environment variable with the authorization key. Pass the key in the `X-API-KEY` header.

## How to run

1. Download database from ip66.dev
2. Mount database as `/ip66.mmdb` inside Docker container
3. Optional: download city database (e.g. GeoLite2-City) and mount as `/city.mmdb`
4. App listens on port 8080.