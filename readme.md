# sloth-job-notifier

Utility executable to notify Sloth about batch job completion.

## Usage

```bash
SLOTH_JOB_NOTIFIER_PROJECT=    # GCP project hosting the Pub/Sub topic
SLOTH_JOB_NOTIFIER_TOPIC=      # Pub/Sub topic name
SLOTH_JOB_NOTIFIER_NAMESPACE=  # Target Sloth namespace

sloth-job-notifier $job_id     # Job ID to notify Sloth about
```
