//init-mongo.sh
#!/bin/bash
set -e

mongo <<EOF
db = db.getSiblingDB('off');
db.createCollection('products');
EOF
