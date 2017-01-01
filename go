#!/bin/bash
docker-compose run app go ${@+"$@"}
