vegeta attack -targets="healthcheck.txt" -duration=30s -rate=1000/s -output="healthcheck.gob"
vegeta report "healthcheck.gob"