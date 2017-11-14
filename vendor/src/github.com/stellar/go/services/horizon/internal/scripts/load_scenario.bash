#! /usr/bin/env bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
SCENARIO=$1
CORE_SQL=$DIR/../test/scenarios/$SCENARIO-core.sql
HORIZON_SQL=$DIR/../test/scenarios/$SCENARIO-horizon.sql

echo "psql $STELLAR_CORE_DATABASE_URL < $CORE_SQL" 
psql $STELLAR_CORE_DATABASE_URL < $CORE_SQL 
echo "psql $DATABASE_URL < $HORIZON_SQL"
psql $DATABASE_URL < $HORIZON_SQL 
