

# Log resources:

### ✅ **Categories of Logs & Where to Get Them**

#### 1. **Web Servers**

* **Apache Access/Error Logs**

  * Format: Common Log Format (CLF)
  * Download from: [Apache log examples](https://github.com/elastic/examples/tree/master/Common%20Data%20Formats/apache_logs)
* **Nginx Access Logs**

  * Format: CLF or JSON
  * Get from: [Sample Logs - Elastic](https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-module-nginx.html)

---

#### 2. **System Logs**

* **Linux Syslog / systemd / auth.log / dmesg**

  * Security, startup, kernel messages, etc.
  * Available via: `/var/log/` on most Linux systems or:
  * Online: [Loghub dataset](https://github.com/logpai/loghub)

---

#### 3. **Application Logs**

* **Python, Node.js, Java apps using structured logging**

  * Log level-based (`INFO`, `ERROR`, etc.)
  * Generate via scripts or use [Loghub app logs](https://github.com/logpai/loghub)

---

#### 4. **Security & Intrusion Detection Logs**

* **SSH brute force, failed logins, suspicious IPs**
* Datasets:

  * [UNSW-NB15 dataset](https://research.unsw.edu.au/projects/unsw-nb15-dataset)
  * [CICIDS logs](https://www.unb.ca/cic/datasets/ids.html)

---

#### 5. **Firewall / Network Logs**

* **Cisco ASA, pfSense, iptables logs**

  * Often in syslog format
  * Example: [Graylog log samples](https://github.com/Graylog2/graylog-guide-syslog-contentpack)

---

#### 6. **Cloud Service Logs**

* **AWS CloudTrail, Azure Monitor, GCP logs**

  * JSON structured logs, API call records
  * Sample logs:

    * AWS: [CloudTrail samples](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-event-reference-record-contents.html)
    * GCP: [GCP Logging Sample](https://cloud.google.com/logging/docs/view/logs-viewer-interface)

---

#### 7. **Database Logs**

* **MySQL, PostgreSQL, MongoDB logs**

  * Includes slow queries, errors, access
  * Sample PostgreSQL logs: [PostgreSQL wiki](https://wiki.postgresql.org/wiki/Logging_Difficulties)
  * Generate via local DB setup with query load scripts

---

#### 8. **Docker & Kubernetes Logs**

* **Container stdout/stderr, K8s pod events**

  * JSON log format or kubectl output
  * Simulate or use public examples:

    * [Docker log samples](https://docs.docker.com/config/containers/logging/)
    * [K8s event logs](https://github.com/kubernetes/kubernetes/issues/33039)

