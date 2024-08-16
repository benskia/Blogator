#!/bin/bash

# HTTPie is used to send test requests.

echo "----- Checking Readiness -----"
http GET localhost:8080/v1/healthz

echo "----- Checking Error Response -----"
curl -v http://localhost:8080/v1/err
http GET localhost:8080/v1/err

echo "----- Creating User -----"
http POST localhost:8080/v1/users name="test"

echo "----- Getting User : Fails : Missing Key -----"
http GET localhost:8080/v1/users

echo "----- Getting User : Fails : Malformed Key -----"
http GET localhost:8080/v1/users Authorization:'Oops 123'

echo "----- Getting User : Succeeds -----"
http GET localhost:8080/v1/users Authorization:'ApiKey 06fc0ecfee4acf2acd5d43c1ecf6dc3d31a3ae8554c1ab0f345670dd06305031'

echo "----- Creating Feed : Fails -----"
http POST localhost:8080/v1/feeds Authorization:'ApiKey invalidkey' name='The Boot.dev Blog' url='https://blog.boot.dev/index.xml'

echo "----- Creating Feed : Succeeds -----"
http POST localhost:8080/v1/feeds Authorization:'ApiKey 06fc0ecfee4acf2acd5d43c1ecf6dc3d31a3ae8554c1ab0f345670dd06305031' name='The Boot.dev Blog' url='https://blog.boot.dev/index.xml'

echo "----- Getting Feeds -----"
http GET localhost:8080/v1/feeds

echo "----- Following Feed -----"
http POST localhost:8080/v1/feed_follows Authorization:'ApiKey 06fc0ecfee4acf2acd5d43c1ecf6dc3d31a3ae8554c1ab0f345670dd06305031' feed_id='ccffd0bc-db78-4de1-9c2f-083d2adb516a'

echo "----- Unfollowing Feed -----"
http DELETE localhost:8080/v1/feed_follows/{3d7b9028-206e-4bdd-88f1-961261fa7a56} Authorization:'ApiKey 06fc0ecfee4acf2acd5d43c1ecf6dc3d31a3ae8554c1ab0f345670dd06305031'

echo "----- Getting Feed Follows -----"
http GET localhost:8080/v1/feed_follows Authorization:'ApiKey 06fc0ecfee4acf2acd5d43c1ecf6dc3d31a3ae8554c1ab0f345670dd06305031'
