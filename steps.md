4. Project-Vendor Matching

Endpoint: /projects/{id}/Matchs/rebuild
Logic:

Vendor’s countries_supported must include project’s country.

Vendor’s services_offered ∩ project’s services_needed must not be empty.

Score = services_overlap * 2 + rating + SLA_weight.

Save to Matchs with upsert logic (avoid duplicates).



5. Analytics & Cross-DB Queries

Endpoint: /analytics/top-vendors

From MySQL → top 3 vendors per country in last 30 days.

From MongoDB → count of documents per country.

Combine results in one response.



6: Analytics

Add endpoint /analytics/top-vendors.

In use case:

Query MySQL: top 3 vendors per country in last 30 days.

Query Mongo: count of docs per country.

Combine into single response.
👉 Test: Call API → see analytics JSON.