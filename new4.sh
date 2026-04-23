# Clean up any old streams, ignoring errors if they don't exist
nats stream rm test-stream -f --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw
nats stream rm DISPATCH_STREAM -f --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw
nats stream rm FEEDBACK_STREAM -f --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw
nats stream rm WATCHDOG_STREAM -f --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw

# --- Re-create all necessary streams ---

echo "Creating 'test-stream' for publisher/subscriber..."
nats stream add test-stream --subjects "foo.test" --retention limits --max-age 24h --storage file --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw

echo "Creating 'DISPATCH_STREAM' for hub..."
nats stream add DISPATCH_STREAM --subjects "dispatch.>" --retention limits --max-age 48h --storage file --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw

echo "Creating 'FEEDBACK_STREAM' for leaf_app..."
nats stream add FEEDBACK_STREAM --subjects "feedback.>" --retention limits --max-age 720h --storage file --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw

echo "Creating 'WATCHDOG_STREAM' for leaf_app..."
nats stream add WATCHDOG_STREAM --subjects "watchdog.>" --retention limits --max-age 1h --storage file --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw