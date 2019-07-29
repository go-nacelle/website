#!/bin/sh -ex

trap "{ rm -f .s3cfg; }" EXIT

cat << EOF > .s3cfg
[default]
access_key = ${ACCESS_KEY}
secret_key = ${SECRET_KEY}
host_base = sfo2.digitaloceanspaces.com
host_bucket = %(bucket)s.sfo2.digitaloceanspaces.com
EOF

s3cmd --config .s3cfg put public s3://laniakea/nacelle/ --acl-public --recursive
