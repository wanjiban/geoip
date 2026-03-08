mod config.json to reduce size of geoip.dat.form 18.9m to 0.2m
      
      "args": {
        "outputDir": "./output",
        "outputName": "geoip.dat"
        "overwriteList": [
          "cn",
          "private",
          "cloudflare",
          "cloudfront",
          "facebook",
          "fastly",
          "google",
          "netflix",
          "telegram",
          "twitter"
        ]
      }
    },
    {
