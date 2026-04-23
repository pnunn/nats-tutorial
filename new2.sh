# Delete existing streams to ensure a clean slate
nats stream rm DISPATCH_STREAM -f --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw
nats stream rm FEEDBACK_STREAM -f --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw
nats stream rm WATCHDOG_STREAM -f --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw

# Re-create them with the correct subjects
nats stream add DISPATCH_STREAM --subjects "dispatch.>" --retention limits --max-age 48h --storage file --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw
nats stream add FEEDBACK_STREAM --subjects "feedback.>" --retention limits --max-age 720h --storage file --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw
nats stream add WATCHDOG_STREAM --subjects "watchdog.>" --retention limits --max-age 1h --storage file --server tls://natsgw.marketdispatch.com.au:4222 --user clientuser --password clIent2026Pw