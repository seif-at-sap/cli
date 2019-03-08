package helpers

import (
	"net/http"

	"github.com/onsi/gomega/ghttp"
)

var fiftyOneOrgJSONPage1 string = `{
	"next_url": "/v2/organizations?order-by=name&order-direction=asc&page=2&results-per-page=50",
	"prev_url": null,
	"resources": [
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/fb1f660d-c87d-4752-97df-92063929b442/app_events",
		  "auditors_url": "/v2/organizations/fb1f660d-c87d-4752-97df-92063929b442/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/fb1f660d-c87d-4752-97df-92063929b442/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/fb1f660d-c87d-4752-97df-92063929b442/domains",
		  "managers_url": "/v2/organizations/fb1f660d-c87d-4752-97df-92063929b442/managers",
		  "name": "org1",
		  "private_domains_url": "/v2/organizations/fb1f660d-c87d-4752-97df-92063929b442/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/fb1f660d-c87d-4752-97df-92063929b442/space_quota_definitions",
		  "spaces_url": "/v2/organizations/fb1f660d-c87d-4752-97df-92063929b442/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/fb1f660d-c87d-4752-97df-92063929b442/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:36:53Z",
		  "guid": "fb1f660d-c87d-4752-97df-92063929b442",
		  "updated_at": "2019-03-05T01:36:53Z",
		  "url": "/v2/organizations/fb1f660d-c87d-4752-97df-92063929b442"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/6f30e06d-360e-4cd7-9849-01f28109bc37/app_events",
		  "auditors_url": "/v2/organizations/6f30e06d-360e-4cd7-9849-01f28109bc37/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/6f30e06d-360e-4cd7-9849-01f28109bc37/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/6f30e06d-360e-4cd7-9849-01f28109bc37/domains",
		  "managers_url": "/v2/organizations/6f30e06d-360e-4cd7-9849-01f28109bc37/managers",
		  "name": "org10",
		  "private_domains_url": "/v2/organizations/6f30e06d-360e-4cd7-9849-01f28109bc37/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/6f30e06d-360e-4cd7-9849-01f28109bc37/space_quota_definitions",
		  "spaces_url": "/v2/organizations/6f30e06d-360e-4cd7-9849-01f28109bc37/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/6f30e06d-360e-4cd7-9849-01f28109bc37/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:36:59Z",
		  "guid": "6f30e06d-360e-4cd7-9849-01f28109bc37",
		  "updated_at": "2019-03-05T01:36:59Z",
		  "url": "/v2/organizations/6f30e06d-360e-4cd7-9849-01f28109bc37"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/459b4d0d-5cf1-46e9-8bd1-b61a4a7266f1/app_events",
		  "auditors_url": "/v2/organizations/459b4d0d-5cf1-46e9-8bd1-b61a4a7266f1/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/459b4d0d-5cf1-46e9-8bd1-b61a4a7266f1/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/459b4d0d-5cf1-46e9-8bd1-b61a4a7266f1/domains",
		  "managers_url": "/v2/organizations/459b4d0d-5cf1-46e9-8bd1-b61a4a7266f1/managers",
		  "name": "org11",
		  "private_domains_url": "/v2/organizations/459b4d0d-5cf1-46e9-8bd1-b61a4a7266f1/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/459b4d0d-5cf1-46e9-8bd1-b61a4a7266f1/space_quota_definitions",
		  "spaces_url": "/v2/organizations/459b4d0d-5cf1-46e9-8bd1-b61a4a7266f1/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/459b4d0d-5cf1-46e9-8bd1-b61a4a7266f1/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:36:59Z",
		  "guid": "459b4d0d-5cf1-46e9-8bd1-b61a4a7266f1",
		  "updated_at": "2019-03-05T01:36:59Z",
		  "url": "/v2/organizations/459b4d0d-5cf1-46e9-8bd1-b61a4a7266f1"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/49f19931-d395-4140-870a-1a26b6ea81e4/app_events",
		  "auditors_url": "/v2/organizations/49f19931-d395-4140-870a-1a26b6ea81e4/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/49f19931-d395-4140-870a-1a26b6ea81e4/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/49f19931-d395-4140-870a-1a26b6ea81e4/domains",
		  "managers_url": "/v2/organizations/49f19931-d395-4140-870a-1a26b6ea81e4/managers",
		  "name": "org12",
		  "private_domains_url": "/v2/organizations/49f19931-d395-4140-870a-1a26b6ea81e4/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/49f19931-d395-4140-870a-1a26b6ea81e4/space_quota_definitions",
		  "spaces_url": "/v2/organizations/49f19931-d395-4140-870a-1a26b6ea81e4/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/49f19931-d395-4140-870a-1a26b6ea81e4/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:00Z",
		  "guid": "49f19931-d395-4140-870a-1a26b6ea81e4",
		  "updated_at": "2019-03-05T01:37:00Z",
		  "url": "/v2/organizations/49f19931-d395-4140-870a-1a26b6ea81e4"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/160faf66-5da3-45cc-b8ee-2501fa56b464/app_events",
		  "auditors_url": "/v2/organizations/160faf66-5da3-45cc-b8ee-2501fa56b464/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/160faf66-5da3-45cc-b8ee-2501fa56b464/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/160faf66-5da3-45cc-b8ee-2501fa56b464/domains",
		  "managers_url": "/v2/organizations/160faf66-5da3-45cc-b8ee-2501fa56b464/managers",
		  "name": "org13",
		  "private_domains_url": "/v2/organizations/160faf66-5da3-45cc-b8ee-2501fa56b464/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/160faf66-5da3-45cc-b8ee-2501fa56b464/space_quota_definitions",
		  "spaces_url": "/v2/organizations/160faf66-5da3-45cc-b8ee-2501fa56b464/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/160faf66-5da3-45cc-b8ee-2501fa56b464/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:00Z",
		  "guid": "160faf66-5da3-45cc-b8ee-2501fa56b464",
		  "updated_at": "2019-03-05T01:37:00Z",
		  "url": "/v2/organizations/160faf66-5da3-45cc-b8ee-2501fa56b464"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/b31b2ffb-8194-40a1-b308-6e11ccf48fac/app_events",
		  "auditors_url": "/v2/organizations/b31b2ffb-8194-40a1-b308-6e11ccf48fac/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/b31b2ffb-8194-40a1-b308-6e11ccf48fac/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/b31b2ffb-8194-40a1-b308-6e11ccf48fac/domains",
		  "managers_url": "/v2/organizations/b31b2ffb-8194-40a1-b308-6e11ccf48fac/managers",
		  "name": "org14",
		  "private_domains_url": "/v2/organizations/b31b2ffb-8194-40a1-b308-6e11ccf48fac/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/b31b2ffb-8194-40a1-b308-6e11ccf48fac/space_quota_definitions",
		  "spaces_url": "/v2/organizations/b31b2ffb-8194-40a1-b308-6e11ccf48fac/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/b31b2ffb-8194-40a1-b308-6e11ccf48fac/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:01Z",
		  "guid": "b31b2ffb-8194-40a1-b308-6e11ccf48fac",
		  "updated_at": "2019-03-05T01:37:01Z",
		  "url": "/v2/organizations/b31b2ffb-8194-40a1-b308-6e11ccf48fac"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/a438f36b-67a5-4432-9744-b33cd7309fff/app_events",
		  "auditors_url": "/v2/organizations/a438f36b-67a5-4432-9744-b33cd7309fff/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/a438f36b-67a5-4432-9744-b33cd7309fff/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/a438f36b-67a5-4432-9744-b33cd7309fff/domains",
		  "managers_url": "/v2/organizations/a438f36b-67a5-4432-9744-b33cd7309fff/managers",
		  "name": "org15",
		  "private_domains_url": "/v2/organizations/a438f36b-67a5-4432-9744-b33cd7309fff/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/a438f36b-67a5-4432-9744-b33cd7309fff/space_quota_definitions",
		  "spaces_url": "/v2/organizations/a438f36b-67a5-4432-9744-b33cd7309fff/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/a438f36b-67a5-4432-9744-b33cd7309fff/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:01Z",
		  "guid": "a438f36b-67a5-4432-9744-b33cd7309fff",
		  "updated_at": "2019-03-05T01:37:01Z",
		  "url": "/v2/organizations/a438f36b-67a5-4432-9744-b33cd7309fff"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/e67ffe78-e479-4fe0-ada7-fcd6ca0238f7/app_events",
		  "auditors_url": "/v2/organizations/e67ffe78-e479-4fe0-ada7-fcd6ca0238f7/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/e67ffe78-e479-4fe0-ada7-fcd6ca0238f7/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/e67ffe78-e479-4fe0-ada7-fcd6ca0238f7/domains",
		  "managers_url": "/v2/organizations/e67ffe78-e479-4fe0-ada7-fcd6ca0238f7/managers",
		  "name": "org16",
		  "private_domains_url": "/v2/organizations/e67ffe78-e479-4fe0-ada7-fcd6ca0238f7/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/e67ffe78-e479-4fe0-ada7-fcd6ca0238f7/space_quota_definitions",
		  "spaces_url": "/v2/organizations/e67ffe78-e479-4fe0-ada7-fcd6ca0238f7/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/e67ffe78-e479-4fe0-ada7-fcd6ca0238f7/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:02Z",
		  "guid": "e67ffe78-e479-4fe0-ada7-fcd6ca0238f7",
		  "updated_at": "2019-03-05T01:37:02Z",
		  "url": "/v2/organizations/e67ffe78-e479-4fe0-ada7-fcd6ca0238f7"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/0ea5bb2d-bcc1-4eef-ae9f-79c4cdbfdb1d/app_events",
		  "auditors_url": "/v2/organizations/0ea5bb2d-bcc1-4eef-ae9f-79c4cdbfdb1d/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/0ea5bb2d-bcc1-4eef-ae9f-79c4cdbfdb1d/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/0ea5bb2d-bcc1-4eef-ae9f-79c4cdbfdb1d/domains",
		  "managers_url": "/v2/organizations/0ea5bb2d-bcc1-4eef-ae9f-79c4cdbfdb1d/managers",
		  "name": "org17",
		  "private_domains_url": "/v2/organizations/0ea5bb2d-bcc1-4eef-ae9f-79c4cdbfdb1d/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/0ea5bb2d-bcc1-4eef-ae9f-79c4cdbfdb1d/space_quota_definitions",
		  "spaces_url": "/v2/organizations/0ea5bb2d-bcc1-4eef-ae9f-79c4cdbfdb1d/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/0ea5bb2d-bcc1-4eef-ae9f-79c4cdbfdb1d/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:03Z",
		  "guid": "0ea5bb2d-bcc1-4eef-ae9f-79c4cdbfdb1d",
		  "updated_at": "2019-03-05T01:37:03Z",
		  "url": "/v2/organizations/0ea5bb2d-bcc1-4eef-ae9f-79c4cdbfdb1d"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/bc9724e3-5c8e-4bea-91d1-6c76dab26f1d/app_events",
		  "auditors_url": "/v2/organizations/bc9724e3-5c8e-4bea-91d1-6c76dab26f1d/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/bc9724e3-5c8e-4bea-91d1-6c76dab26f1d/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/bc9724e3-5c8e-4bea-91d1-6c76dab26f1d/domains",
		  "managers_url": "/v2/organizations/bc9724e3-5c8e-4bea-91d1-6c76dab26f1d/managers",
		  "name": "org18",
		  "private_domains_url": "/v2/organizations/bc9724e3-5c8e-4bea-91d1-6c76dab26f1d/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/bc9724e3-5c8e-4bea-91d1-6c76dab26f1d/space_quota_definitions",
		  "spaces_url": "/v2/organizations/bc9724e3-5c8e-4bea-91d1-6c76dab26f1d/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/bc9724e3-5c8e-4bea-91d1-6c76dab26f1d/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:03Z",
		  "guid": "bc9724e3-5c8e-4bea-91d1-6c76dab26f1d",
		  "updated_at": "2019-03-05T01:37:03Z",
		  "url": "/v2/organizations/bc9724e3-5c8e-4bea-91d1-6c76dab26f1d"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/0db7a902-03ec-474a-97b4-2f9f0b8fc28a/app_events",
		  "auditors_url": "/v2/organizations/0db7a902-03ec-474a-97b4-2f9f0b8fc28a/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/0db7a902-03ec-474a-97b4-2f9f0b8fc28a/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/0db7a902-03ec-474a-97b4-2f9f0b8fc28a/domains",
		  "managers_url": "/v2/organizations/0db7a902-03ec-474a-97b4-2f9f0b8fc28a/managers",
		  "name": "org19",
		  "private_domains_url": "/v2/organizations/0db7a902-03ec-474a-97b4-2f9f0b8fc28a/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/0db7a902-03ec-474a-97b4-2f9f0b8fc28a/space_quota_definitions",
		  "spaces_url": "/v2/organizations/0db7a902-03ec-474a-97b4-2f9f0b8fc28a/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/0db7a902-03ec-474a-97b4-2f9f0b8fc28a/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:04Z",
		  "guid": "0db7a902-03ec-474a-97b4-2f9f0b8fc28a",
		  "updated_at": "2019-03-05T01:37:04Z",
		  "url": "/v2/organizations/0db7a902-03ec-474a-97b4-2f9f0b8fc28a"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/838dd065-b2d5-400c-96d4-d0068a314e3f/app_events",
		  "auditors_url": "/v2/organizations/838dd065-b2d5-400c-96d4-d0068a314e3f/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/838dd065-b2d5-400c-96d4-d0068a314e3f/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/838dd065-b2d5-400c-96d4-d0068a314e3f/domains",
		  "managers_url": "/v2/organizations/838dd065-b2d5-400c-96d4-d0068a314e3f/managers",
		  "name": "org2",
		  "private_domains_url": "/v2/organizations/838dd065-b2d5-400c-96d4-d0068a314e3f/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/838dd065-b2d5-400c-96d4-d0068a314e3f/space_quota_definitions",
		  "spaces_url": "/v2/organizations/838dd065-b2d5-400c-96d4-d0068a314e3f/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/838dd065-b2d5-400c-96d4-d0068a314e3f/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:36:54Z",
		  "guid": "838dd065-b2d5-400c-96d4-d0068a314e3f",
		  "updated_at": "2019-03-05T01:36:54Z",
		  "url": "/v2/organizations/838dd065-b2d5-400c-96d4-d0068a314e3f"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/dce09b3a-fa81-4efb-ba6c-998dfc66729a/app_events",
		  "auditors_url": "/v2/organizations/dce09b3a-fa81-4efb-ba6c-998dfc66729a/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/dce09b3a-fa81-4efb-ba6c-998dfc66729a/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/dce09b3a-fa81-4efb-ba6c-998dfc66729a/domains",
		  "managers_url": "/v2/organizations/dce09b3a-fa81-4efb-ba6c-998dfc66729a/managers",
		  "name": "org20",
		  "private_domains_url": "/v2/organizations/dce09b3a-fa81-4efb-ba6c-998dfc66729a/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/dce09b3a-fa81-4efb-ba6c-998dfc66729a/space_quota_definitions",
		  "spaces_url": "/v2/organizations/dce09b3a-fa81-4efb-ba6c-998dfc66729a/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/dce09b3a-fa81-4efb-ba6c-998dfc66729a/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:04Z",
		  "guid": "dce09b3a-fa81-4efb-ba6c-998dfc66729a",
		  "updated_at": "2019-03-05T01:37:04Z",
		  "url": "/v2/organizations/dce09b3a-fa81-4efb-ba6c-998dfc66729a"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/4cbc8085-b35f-4c4e-952d-3bb8a7aa2130/app_events",
		  "auditors_url": "/v2/organizations/4cbc8085-b35f-4c4e-952d-3bb8a7aa2130/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/4cbc8085-b35f-4c4e-952d-3bb8a7aa2130/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/4cbc8085-b35f-4c4e-952d-3bb8a7aa2130/domains",
		  "managers_url": "/v2/organizations/4cbc8085-b35f-4c4e-952d-3bb8a7aa2130/managers",
		  "name": "org21",
		  "private_domains_url": "/v2/organizations/4cbc8085-b35f-4c4e-952d-3bb8a7aa2130/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/4cbc8085-b35f-4c4e-952d-3bb8a7aa2130/space_quota_definitions",
		  "spaces_url": "/v2/organizations/4cbc8085-b35f-4c4e-952d-3bb8a7aa2130/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/4cbc8085-b35f-4c4e-952d-3bb8a7aa2130/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:05Z",
		  "guid": "4cbc8085-b35f-4c4e-952d-3bb8a7aa2130",
		  "updated_at": "2019-03-05T01:37:05Z",
		  "url": "/v2/organizations/4cbc8085-b35f-4c4e-952d-3bb8a7aa2130"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/55f98f05-cd65-4dd7-9309-c76cd93e9f8c/app_events",
		  "auditors_url": "/v2/organizations/55f98f05-cd65-4dd7-9309-c76cd93e9f8c/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/55f98f05-cd65-4dd7-9309-c76cd93e9f8c/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/55f98f05-cd65-4dd7-9309-c76cd93e9f8c/domains",
		  "managers_url": "/v2/organizations/55f98f05-cd65-4dd7-9309-c76cd93e9f8c/managers",
		  "name": "org22",
		  "private_domains_url": "/v2/organizations/55f98f05-cd65-4dd7-9309-c76cd93e9f8c/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/55f98f05-cd65-4dd7-9309-c76cd93e9f8c/space_quota_definitions",
		  "spaces_url": "/v2/organizations/55f98f05-cd65-4dd7-9309-c76cd93e9f8c/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/55f98f05-cd65-4dd7-9309-c76cd93e9f8c/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:05Z",
		  "guid": "55f98f05-cd65-4dd7-9309-c76cd93e9f8c",
		  "updated_at": "2019-03-05T01:37:05Z",
		  "url": "/v2/organizations/55f98f05-cd65-4dd7-9309-c76cd93e9f8c"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/0018b94d-4bf9-4f7c-b69f-31aa1dab7756/app_events",
		  "auditors_url": "/v2/organizations/0018b94d-4bf9-4f7c-b69f-31aa1dab7756/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/0018b94d-4bf9-4f7c-b69f-31aa1dab7756/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/0018b94d-4bf9-4f7c-b69f-31aa1dab7756/domains",
		  "managers_url": "/v2/organizations/0018b94d-4bf9-4f7c-b69f-31aa1dab7756/managers",
		  "name": "org23",
		  "private_domains_url": "/v2/organizations/0018b94d-4bf9-4f7c-b69f-31aa1dab7756/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/0018b94d-4bf9-4f7c-b69f-31aa1dab7756/space_quota_definitions",
		  "spaces_url": "/v2/organizations/0018b94d-4bf9-4f7c-b69f-31aa1dab7756/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/0018b94d-4bf9-4f7c-b69f-31aa1dab7756/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:06Z",
		  "guid": "0018b94d-4bf9-4f7c-b69f-31aa1dab7756",
		  "updated_at": "2019-03-05T01:37:06Z",
		  "url": "/v2/organizations/0018b94d-4bf9-4f7c-b69f-31aa1dab7756"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/5597c00e-b05d-406b-826b-1c308e070df4/app_events",
		  "auditors_url": "/v2/organizations/5597c00e-b05d-406b-826b-1c308e070df4/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/5597c00e-b05d-406b-826b-1c308e070df4/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/5597c00e-b05d-406b-826b-1c308e070df4/domains",
		  "managers_url": "/v2/organizations/5597c00e-b05d-406b-826b-1c308e070df4/managers",
		  "name": "org24",
		  "private_domains_url": "/v2/organizations/5597c00e-b05d-406b-826b-1c308e070df4/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/5597c00e-b05d-406b-826b-1c308e070df4/space_quota_definitions",
		  "spaces_url": "/v2/organizations/5597c00e-b05d-406b-826b-1c308e070df4/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/5597c00e-b05d-406b-826b-1c308e070df4/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:06Z",
		  "guid": "5597c00e-b05d-406b-826b-1c308e070df4",
		  "updated_at": "2019-03-05T01:37:06Z",
		  "url": "/v2/organizations/5597c00e-b05d-406b-826b-1c308e070df4"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/0cc3f28e-9722-4de5-8f68-c8e7e30710ca/app_events",
		  "auditors_url": "/v2/organizations/0cc3f28e-9722-4de5-8f68-c8e7e30710ca/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/0cc3f28e-9722-4de5-8f68-c8e7e30710ca/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/0cc3f28e-9722-4de5-8f68-c8e7e30710ca/domains",
		  "managers_url": "/v2/organizations/0cc3f28e-9722-4de5-8f68-c8e7e30710ca/managers",
		  "name": "org25",
		  "private_domains_url": "/v2/organizations/0cc3f28e-9722-4de5-8f68-c8e7e30710ca/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/0cc3f28e-9722-4de5-8f68-c8e7e30710ca/space_quota_definitions",
		  "spaces_url": "/v2/organizations/0cc3f28e-9722-4de5-8f68-c8e7e30710ca/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/0cc3f28e-9722-4de5-8f68-c8e7e30710ca/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:07Z",
		  "guid": "0cc3f28e-9722-4de5-8f68-c8e7e30710ca",
		  "updated_at": "2019-03-05T01:37:07Z",
		  "url": "/v2/organizations/0cc3f28e-9722-4de5-8f68-c8e7e30710ca"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/9065288f-6490-41f9-ba6f-02c09b9781c6/app_events",
		  "auditors_url": "/v2/organizations/9065288f-6490-41f9-ba6f-02c09b9781c6/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/9065288f-6490-41f9-ba6f-02c09b9781c6/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/9065288f-6490-41f9-ba6f-02c09b9781c6/domains",
		  "managers_url": "/v2/organizations/9065288f-6490-41f9-ba6f-02c09b9781c6/managers",
		  "name": "org26",
		  "private_domains_url": "/v2/organizations/9065288f-6490-41f9-ba6f-02c09b9781c6/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/9065288f-6490-41f9-ba6f-02c09b9781c6/space_quota_definitions",
		  "spaces_url": "/v2/organizations/9065288f-6490-41f9-ba6f-02c09b9781c6/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/9065288f-6490-41f9-ba6f-02c09b9781c6/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:08Z",
		  "guid": "9065288f-6490-41f9-ba6f-02c09b9781c6",
		  "updated_at": "2019-03-05T01:37:08Z",
		  "url": "/v2/organizations/9065288f-6490-41f9-ba6f-02c09b9781c6"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/573c3d33-74b2-4f9d-bc53-cfbe90f832e4/app_events",
		  "auditors_url": "/v2/organizations/573c3d33-74b2-4f9d-bc53-cfbe90f832e4/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/573c3d33-74b2-4f9d-bc53-cfbe90f832e4/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/573c3d33-74b2-4f9d-bc53-cfbe90f832e4/domains",
		  "managers_url": "/v2/organizations/573c3d33-74b2-4f9d-bc53-cfbe90f832e4/managers",
		  "name": "org27",
		  "private_domains_url": "/v2/organizations/573c3d33-74b2-4f9d-bc53-cfbe90f832e4/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/573c3d33-74b2-4f9d-bc53-cfbe90f832e4/space_quota_definitions",
		  "spaces_url": "/v2/organizations/573c3d33-74b2-4f9d-bc53-cfbe90f832e4/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/573c3d33-74b2-4f9d-bc53-cfbe90f832e4/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:08Z",
		  "guid": "573c3d33-74b2-4f9d-bc53-cfbe90f832e4",
		  "updated_at": "2019-03-05T01:37:08Z",
		  "url": "/v2/organizations/573c3d33-74b2-4f9d-bc53-cfbe90f832e4"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/9dd39bed-e524-4d4b-bfb1-b0dc8bcde9f4/app_events",
		  "auditors_url": "/v2/organizations/9dd39bed-e524-4d4b-bfb1-b0dc8bcde9f4/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/9dd39bed-e524-4d4b-bfb1-b0dc8bcde9f4/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/9dd39bed-e524-4d4b-bfb1-b0dc8bcde9f4/domains",
		  "managers_url": "/v2/organizations/9dd39bed-e524-4d4b-bfb1-b0dc8bcde9f4/managers",
		  "name": "org28",
		  "private_domains_url": "/v2/organizations/9dd39bed-e524-4d4b-bfb1-b0dc8bcde9f4/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/9dd39bed-e524-4d4b-bfb1-b0dc8bcde9f4/space_quota_definitions",
		  "spaces_url": "/v2/organizations/9dd39bed-e524-4d4b-bfb1-b0dc8bcde9f4/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/9dd39bed-e524-4d4b-bfb1-b0dc8bcde9f4/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:09Z",
		  "guid": "9dd39bed-e524-4d4b-bfb1-b0dc8bcde9f4",
		  "updated_at": "2019-03-05T01:37:09Z",
		  "url": "/v2/organizations/9dd39bed-e524-4d4b-bfb1-b0dc8bcde9f4"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/b4dab8b3-32a0-442e-9e28-f8313ad2d49c/app_events",
		  "auditors_url": "/v2/organizations/b4dab8b3-32a0-442e-9e28-f8313ad2d49c/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/b4dab8b3-32a0-442e-9e28-f8313ad2d49c/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/b4dab8b3-32a0-442e-9e28-f8313ad2d49c/domains",
		  "managers_url": "/v2/organizations/b4dab8b3-32a0-442e-9e28-f8313ad2d49c/managers",
		  "name": "org29",
		  "private_domains_url": "/v2/organizations/b4dab8b3-32a0-442e-9e28-f8313ad2d49c/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/b4dab8b3-32a0-442e-9e28-f8313ad2d49c/space_quota_definitions",
		  "spaces_url": "/v2/organizations/b4dab8b3-32a0-442e-9e28-f8313ad2d49c/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/b4dab8b3-32a0-442e-9e28-f8313ad2d49c/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:09Z",
		  "guid": "b4dab8b3-32a0-442e-9e28-f8313ad2d49c",
		  "updated_at": "2019-03-05T01:37:09Z",
		  "url": "/v2/organizations/b4dab8b3-32a0-442e-9e28-f8313ad2d49c"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/39c33a35-6155-4146-9e82-31d1edc9c546/app_events",
		  "auditors_url": "/v2/organizations/39c33a35-6155-4146-9e82-31d1edc9c546/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/39c33a35-6155-4146-9e82-31d1edc9c546/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/39c33a35-6155-4146-9e82-31d1edc9c546/domains",
		  "managers_url": "/v2/organizations/39c33a35-6155-4146-9e82-31d1edc9c546/managers",
		  "name": "org3",
		  "private_domains_url": "/v2/organizations/39c33a35-6155-4146-9e82-31d1edc9c546/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/39c33a35-6155-4146-9e82-31d1edc9c546/space_quota_definitions",
		  "spaces_url": "/v2/organizations/39c33a35-6155-4146-9e82-31d1edc9c546/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/39c33a35-6155-4146-9e82-31d1edc9c546/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:36:55Z",
		  "guid": "39c33a35-6155-4146-9e82-31d1edc9c546",
		  "updated_at": "2019-03-05T01:36:55Z",
		  "url": "/v2/organizations/39c33a35-6155-4146-9e82-31d1edc9c546"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/63a760fc-413d-49dd-9390-9302ab8220f0/app_events",
		  "auditors_url": "/v2/organizations/63a760fc-413d-49dd-9390-9302ab8220f0/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/63a760fc-413d-49dd-9390-9302ab8220f0/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/63a760fc-413d-49dd-9390-9302ab8220f0/domains",
		  "managers_url": "/v2/organizations/63a760fc-413d-49dd-9390-9302ab8220f0/managers",
		  "name": "org30",
		  "private_domains_url": "/v2/organizations/63a760fc-413d-49dd-9390-9302ab8220f0/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/63a760fc-413d-49dd-9390-9302ab8220f0/space_quota_definitions",
		  "spaces_url": "/v2/organizations/63a760fc-413d-49dd-9390-9302ab8220f0/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/63a760fc-413d-49dd-9390-9302ab8220f0/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:10Z",
		  "guid": "63a760fc-413d-49dd-9390-9302ab8220f0",
		  "updated_at": "2019-03-05T01:37:10Z",
		  "url": "/v2/organizations/63a760fc-413d-49dd-9390-9302ab8220f0"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/db6ac1a4-6740-4826-874e-4fc4b479007f/app_events",
		  "auditors_url": "/v2/organizations/db6ac1a4-6740-4826-874e-4fc4b479007f/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/db6ac1a4-6740-4826-874e-4fc4b479007f/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/db6ac1a4-6740-4826-874e-4fc4b479007f/domains",
		  "managers_url": "/v2/organizations/db6ac1a4-6740-4826-874e-4fc4b479007f/managers",
		  "name": "org31",
		  "private_domains_url": "/v2/organizations/db6ac1a4-6740-4826-874e-4fc4b479007f/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/db6ac1a4-6740-4826-874e-4fc4b479007f/space_quota_definitions",
		  "spaces_url": "/v2/organizations/db6ac1a4-6740-4826-874e-4fc4b479007f/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/db6ac1a4-6740-4826-874e-4fc4b479007f/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:10Z",
		  "guid": "db6ac1a4-6740-4826-874e-4fc4b479007f",
		  "updated_at": "2019-03-05T01:37:10Z",
		  "url": "/v2/organizations/db6ac1a4-6740-4826-874e-4fc4b479007f"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/31481f15-af48-4b85-b655-8453dc407221/app_events",
		  "auditors_url": "/v2/organizations/31481f15-af48-4b85-b655-8453dc407221/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/31481f15-af48-4b85-b655-8453dc407221/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/31481f15-af48-4b85-b655-8453dc407221/domains",
		  "managers_url": "/v2/organizations/31481f15-af48-4b85-b655-8453dc407221/managers",
		  "name": "org32",
		  "private_domains_url": "/v2/organizations/31481f15-af48-4b85-b655-8453dc407221/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/31481f15-af48-4b85-b655-8453dc407221/space_quota_definitions",
		  "spaces_url": "/v2/organizations/31481f15-af48-4b85-b655-8453dc407221/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/31481f15-af48-4b85-b655-8453dc407221/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:11Z",
		  "guid": "31481f15-af48-4b85-b655-8453dc407221",
		  "updated_at": "2019-03-05T01:37:11Z",
		  "url": "/v2/organizations/31481f15-af48-4b85-b655-8453dc407221"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/992eb9c1-16bb-4bc6-9ca1-df4026607262/app_events",
		  "auditors_url": "/v2/organizations/992eb9c1-16bb-4bc6-9ca1-df4026607262/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/992eb9c1-16bb-4bc6-9ca1-df4026607262/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/992eb9c1-16bb-4bc6-9ca1-df4026607262/domains",
		  "managers_url": "/v2/organizations/992eb9c1-16bb-4bc6-9ca1-df4026607262/managers",
		  "name": "org33",
		  "private_domains_url": "/v2/organizations/992eb9c1-16bb-4bc6-9ca1-df4026607262/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/992eb9c1-16bb-4bc6-9ca1-df4026607262/space_quota_definitions",
		  "spaces_url": "/v2/organizations/992eb9c1-16bb-4bc6-9ca1-df4026607262/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/992eb9c1-16bb-4bc6-9ca1-df4026607262/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:12Z",
		  "guid": "992eb9c1-16bb-4bc6-9ca1-df4026607262",
		  "updated_at": "2019-03-05T01:37:12Z",
		  "url": "/v2/organizations/992eb9c1-16bb-4bc6-9ca1-df4026607262"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/398d9998-bcb0-409f-8008-6eecd70672c0/app_events",
		  "auditors_url": "/v2/organizations/398d9998-bcb0-409f-8008-6eecd70672c0/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/398d9998-bcb0-409f-8008-6eecd70672c0/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/398d9998-bcb0-409f-8008-6eecd70672c0/domains",
		  "managers_url": "/v2/organizations/398d9998-bcb0-409f-8008-6eecd70672c0/managers",
		  "name": "org34",
		  "private_domains_url": "/v2/organizations/398d9998-bcb0-409f-8008-6eecd70672c0/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/398d9998-bcb0-409f-8008-6eecd70672c0/space_quota_definitions",
		  "spaces_url": "/v2/organizations/398d9998-bcb0-409f-8008-6eecd70672c0/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/398d9998-bcb0-409f-8008-6eecd70672c0/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:12Z",
		  "guid": "398d9998-bcb0-409f-8008-6eecd70672c0",
		  "updated_at": "2019-03-05T01:37:12Z",
		  "url": "/v2/organizations/398d9998-bcb0-409f-8008-6eecd70672c0"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/7e66e6f6-07a1-4226-b845-cd6a1ab1b49a/app_events",
		  "auditors_url": "/v2/organizations/7e66e6f6-07a1-4226-b845-cd6a1ab1b49a/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/7e66e6f6-07a1-4226-b845-cd6a1ab1b49a/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/7e66e6f6-07a1-4226-b845-cd6a1ab1b49a/domains",
		  "managers_url": "/v2/organizations/7e66e6f6-07a1-4226-b845-cd6a1ab1b49a/managers",
		  "name": "org35",
		  "private_domains_url": "/v2/organizations/7e66e6f6-07a1-4226-b845-cd6a1ab1b49a/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/7e66e6f6-07a1-4226-b845-cd6a1ab1b49a/space_quota_definitions",
		  "spaces_url": "/v2/organizations/7e66e6f6-07a1-4226-b845-cd6a1ab1b49a/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/7e66e6f6-07a1-4226-b845-cd6a1ab1b49a/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:13Z",
		  "guid": "7e66e6f6-07a1-4226-b845-cd6a1ab1b49a",
		  "updated_at": "2019-03-05T01:37:13Z",
		  "url": "/v2/organizations/7e66e6f6-07a1-4226-b845-cd6a1ab1b49a"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/c6da7ab4-004f-43d4-8737-a8f48452babb/app_events",
		  "auditors_url": "/v2/organizations/c6da7ab4-004f-43d4-8737-a8f48452babb/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/c6da7ab4-004f-43d4-8737-a8f48452babb/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/c6da7ab4-004f-43d4-8737-a8f48452babb/domains",
		  "managers_url": "/v2/organizations/c6da7ab4-004f-43d4-8737-a8f48452babb/managers",
		  "name": "org36",
		  "private_domains_url": "/v2/organizations/c6da7ab4-004f-43d4-8737-a8f48452babb/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/c6da7ab4-004f-43d4-8737-a8f48452babb/space_quota_definitions",
		  "spaces_url": "/v2/organizations/c6da7ab4-004f-43d4-8737-a8f48452babb/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/c6da7ab4-004f-43d4-8737-a8f48452babb/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:13Z",
		  "guid": "c6da7ab4-004f-43d4-8737-a8f48452babb",
		  "updated_at": "2019-03-05T01:37:13Z",
		  "url": "/v2/organizations/c6da7ab4-004f-43d4-8737-a8f48452babb"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/92625b12-fd5d-4730-b6ae-19ee9f3c38b2/app_events",
		  "auditors_url": "/v2/organizations/92625b12-fd5d-4730-b6ae-19ee9f3c38b2/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/92625b12-fd5d-4730-b6ae-19ee9f3c38b2/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/92625b12-fd5d-4730-b6ae-19ee9f3c38b2/domains",
		  "managers_url": "/v2/organizations/92625b12-fd5d-4730-b6ae-19ee9f3c38b2/managers",
		  "name": "org37",
		  "private_domains_url": "/v2/organizations/92625b12-fd5d-4730-b6ae-19ee9f3c38b2/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/92625b12-fd5d-4730-b6ae-19ee9f3c38b2/space_quota_definitions",
		  "spaces_url": "/v2/organizations/92625b12-fd5d-4730-b6ae-19ee9f3c38b2/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/92625b12-fd5d-4730-b6ae-19ee9f3c38b2/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:14Z",
		  "guid": "92625b12-fd5d-4730-b6ae-19ee9f3c38b2",
		  "updated_at": "2019-03-05T01:37:14Z",
		  "url": "/v2/organizations/92625b12-fd5d-4730-b6ae-19ee9f3c38b2"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/27e1593c-6ce8-475d-8dab-cc5224b02ade/app_events",
		  "auditors_url": "/v2/organizations/27e1593c-6ce8-475d-8dab-cc5224b02ade/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/27e1593c-6ce8-475d-8dab-cc5224b02ade/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/27e1593c-6ce8-475d-8dab-cc5224b02ade/domains",
		  "managers_url": "/v2/organizations/27e1593c-6ce8-475d-8dab-cc5224b02ade/managers",
		  "name": "org38",
		  "private_domains_url": "/v2/organizations/27e1593c-6ce8-475d-8dab-cc5224b02ade/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/27e1593c-6ce8-475d-8dab-cc5224b02ade/space_quota_definitions",
		  "spaces_url": "/v2/organizations/27e1593c-6ce8-475d-8dab-cc5224b02ade/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/27e1593c-6ce8-475d-8dab-cc5224b02ade/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:14Z",
		  "guid": "27e1593c-6ce8-475d-8dab-cc5224b02ade",
		  "updated_at": "2019-03-05T01:37:14Z",
		  "url": "/v2/organizations/27e1593c-6ce8-475d-8dab-cc5224b02ade"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/9dff2351-9d82-4b14-b455-d12d3b220545/app_events",
		  "auditors_url": "/v2/organizations/9dff2351-9d82-4b14-b455-d12d3b220545/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/9dff2351-9d82-4b14-b455-d12d3b220545/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/9dff2351-9d82-4b14-b455-d12d3b220545/domains",
		  "managers_url": "/v2/organizations/9dff2351-9d82-4b14-b455-d12d3b220545/managers",
		  "name": "org39",
		  "private_domains_url": "/v2/organizations/9dff2351-9d82-4b14-b455-d12d3b220545/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/9dff2351-9d82-4b14-b455-d12d3b220545/space_quota_definitions",
		  "spaces_url": "/v2/organizations/9dff2351-9d82-4b14-b455-d12d3b220545/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/9dff2351-9d82-4b14-b455-d12d3b220545/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:15Z",
		  "guid": "9dff2351-9d82-4b14-b455-d12d3b220545",
		  "updated_at": "2019-03-05T01:37:15Z",
		  "url": "/v2/organizations/9dff2351-9d82-4b14-b455-d12d3b220545"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/cf527f93-c359-47b9-9dc8-cee86ca3b338/app_events",
		  "auditors_url": "/v2/organizations/cf527f93-c359-47b9-9dc8-cee86ca3b338/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/cf527f93-c359-47b9-9dc8-cee86ca3b338/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/cf527f93-c359-47b9-9dc8-cee86ca3b338/domains",
		  "managers_url": "/v2/organizations/cf527f93-c359-47b9-9dc8-cee86ca3b338/managers",
		  "name": "org4",
		  "private_domains_url": "/v2/organizations/cf527f93-c359-47b9-9dc8-cee86ca3b338/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/cf527f93-c359-47b9-9dc8-cee86ca3b338/space_quota_definitions",
		  "spaces_url": "/v2/organizations/cf527f93-c359-47b9-9dc8-cee86ca3b338/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/cf527f93-c359-47b9-9dc8-cee86ca3b338/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:36:55Z",
		  "guid": "cf527f93-c359-47b9-9dc8-cee86ca3b338",
		  "updated_at": "2019-03-05T01:36:55Z",
		  "url": "/v2/organizations/cf527f93-c359-47b9-9dc8-cee86ca3b338"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/58fba656-c1e7-4d85-a2d8-9fc4d34d5c38/app_events",
		  "auditors_url": "/v2/organizations/58fba656-c1e7-4d85-a2d8-9fc4d34d5c38/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/58fba656-c1e7-4d85-a2d8-9fc4d34d5c38/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/58fba656-c1e7-4d85-a2d8-9fc4d34d5c38/domains",
		  "managers_url": "/v2/organizations/58fba656-c1e7-4d85-a2d8-9fc4d34d5c38/managers",
		  "name": "org40",
		  "private_domains_url": "/v2/organizations/58fba656-c1e7-4d85-a2d8-9fc4d34d5c38/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/58fba656-c1e7-4d85-a2d8-9fc4d34d5c38/space_quota_definitions",
		  "spaces_url": "/v2/organizations/58fba656-c1e7-4d85-a2d8-9fc4d34d5c38/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/58fba656-c1e7-4d85-a2d8-9fc4d34d5c38/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:15Z",
		  "guid": "58fba656-c1e7-4d85-a2d8-9fc4d34d5c38",
		  "updated_at": "2019-03-05T01:37:15Z",
		  "url": "/v2/organizations/58fba656-c1e7-4d85-a2d8-9fc4d34d5c38"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/8ab80bbb-347c-41d5-9faf-94ea30bfca8a/app_events",
		  "auditors_url": "/v2/organizations/8ab80bbb-347c-41d5-9faf-94ea30bfca8a/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/8ab80bbb-347c-41d5-9faf-94ea30bfca8a/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/8ab80bbb-347c-41d5-9faf-94ea30bfca8a/domains",
		  "managers_url": "/v2/organizations/8ab80bbb-347c-41d5-9faf-94ea30bfca8a/managers",
		  "name": "org41",
		  "private_domains_url": "/v2/organizations/8ab80bbb-347c-41d5-9faf-94ea30bfca8a/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/8ab80bbb-347c-41d5-9faf-94ea30bfca8a/space_quota_definitions",
		  "spaces_url": "/v2/organizations/8ab80bbb-347c-41d5-9faf-94ea30bfca8a/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/8ab80bbb-347c-41d5-9faf-94ea30bfca8a/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:16Z",
		  "guid": "8ab80bbb-347c-41d5-9faf-94ea30bfca8a",
		  "updated_at": "2019-03-05T01:37:16Z",
		  "url": "/v2/organizations/8ab80bbb-347c-41d5-9faf-94ea30bfca8a"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/623cbffd-7194-46a8-b419-555470705a1f/app_events",
		  "auditors_url": "/v2/organizations/623cbffd-7194-46a8-b419-555470705a1f/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/623cbffd-7194-46a8-b419-555470705a1f/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/623cbffd-7194-46a8-b419-555470705a1f/domains",
		  "managers_url": "/v2/organizations/623cbffd-7194-46a8-b419-555470705a1f/managers",
		  "name": "org42",
		  "private_domains_url": "/v2/organizations/623cbffd-7194-46a8-b419-555470705a1f/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/623cbffd-7194-46a8-b419-555470705a1f/space_quota_definitions",
		  "spaces_url": "/v2/organizations/623cbffd-7194-46a8-b419-555470705a1f/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/623cbffd-7194-46a8-b419-555470705a1f/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:17Z",
		  "guid": "623cbffd-7194-46a8-b419-555470705a1f",
		  "updated_at": "2019-03-05T01:37:17Z",
		  "url": "/v2/organizations/623cbffd-7194-46a8-b419-555470705a1f"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/71c9ea51-929e-438c-b5d0-027e500310ba/app_events",
		  "auditors_url": "/v2/organizations/71c9ea51-929e-438c-b5d0-027e500310ba/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/71c9ea51-929e-438c-b5d0-027e500310ba/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/71c9ea51-929e-438c-b5d0-027e500310ba/domains",
		  "managers_url": "/v2/organizations/71c9ea51-929e-438c-b5d0-027e500310ba/managers",
		  "name": "org43",
		  "private_domains_url": "/v2/organizations/71c9ea51-929e-438c-b5d0-027e500310ba/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/71c9ea51-929e-438c-b5d0-027e500310ba/space_quota_definitions",
		  "spaces_url": "/v2/organizations/71c9ea51-929e-438c-b5d0-027e500310ba/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/71c9ea51-929e-438c-b5d0-027e500310ba/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:17Z",
		  "guid": "71c9ea51-929e-438c-b5d0-027e500310ba",
		  "updated_at": "2019-03-05T01:37:17Z",
		  "url": "/v2/organizations/71c9ea51-929e-438c-b5d0-027e500310ba"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/c0d65cc3-f762-4355-b1c7-51a80f7cffe2/app_events",
		  "auditors_url": "/v2/organizations/c0d65cc3-f762-4355-b1c7-51a80f7cffe2/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/c0d65cc3-f762-4355-b1c7-51a80f7cffe2/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/c0d65cc3-f762-4355-b1c7-51a80f7cffe2/domains",
		  "managers_url": "/v2/organizations/c0d65cc3-f762-4355-b1c7-51a80f7cffe2/managers",
		  "name": "org44",
		  "private_domains_url": "/v2/organizations/c0d65cc3-f762-4355-b1c7-51a80f7cffe2/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/c0d65cc3-f762-4355-b1c7-51a80f7cffe2/space_quota_definitions",
		  "spaces_url": "/v2/organizations/c0d65cc3-f762-4355-b1c7-51a80f7cffe2/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/c0d65cc3-f762-4355-b1c7-51a80f7cffe2/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:18Z",
		  "guid": "c0d65cc3-f762-4355-b1c7-51a80f7cffe2",
		  "updated_at": "2019-03-05T01:37:18Z",
		  "url": "/v2/organizations/c0d65cc3-f762-4355-b1c7-51a80f7cffe2"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/e1c9ec7a-3c33-4553-893a-d9b4eb472e95/app_events",
		  "auditors_url": "/v2/organizations/e1c9ec7a-3c33-4553-893a-d9b4eb472e95/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/e1c9ec7a-3c33-4553-893a-d9b4eb472e95/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/e1c9ec7a-3c33-4553-893a-d9b4eb472e95/domains",
		  "managers_url": "/v2/organizations/e1c9ec7a-3c33-4553-893a-d9b4eb472e95/managers",
		  "name": "org45",
		  "private_domains_url": "/v2/organizations/e1c9ec7a-3c33-4553-893a-d9b4eb472e95/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/e1c9ec7a-3c33-4553-893a-d9b4eb472e95/space_quota_definitions",
		  "spaces_url": "/v2/organizations/e1c9ec7a-3c33-4553-893a-d9b4eb472e95/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/e1c9ec7a-3c33-4553-893a-d9b4eb472e95/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:18Z",
		  "guid": "e1c9ec7a-3c33-4553-893a-d9b4eb472e95",
		  "updated_at": "2019-03-05T01:37:18Z",
		  "url": "/v2/organizations/e1c9ec7a-3c33-4553-893a-d9b4eb472e95"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/d256f998-6321-4e96-a3e0-213230fe305b/app_events",
		  "auditors_url": "/v2/organizations/d256f998-6321-4e96-a3e0-213230fe305b/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/d256f998-6321-4e96-a3e0-213230fe305b/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/d256f998-6321-4e96-a3e0-213230fe305b/domains",
		  "managers_url": "/v2/organizations/d256f998-6321-4e96-a3e0-213230fe305b/managers",
		  "name": "org46",
		  "private_domains_url": "/v2/organizations/d256f998-6321-4e96-a3e0-213230fe305b/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/d256f998-6321-4e96-a3e0-213230fe305b/space_quota_definitions",
		  "spaces_url": "/v2/organizations/d256f998-6321-4e96-a3e0-213230fe305b/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/d256f998-6321-4e96-a3e0-213230fe305b/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:19Z",
		  "guid": "d256f998-6321-4e96-a3e0-213230fe305b",
		  "updated_at": "2019-03-05T01:37:19Z",
		  "url": "/v2/organizations/d256f998-6321-4e96-a3e0-213230fe305b"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/5eb16637-a65e-4236-9512-583317cf0387/app_events",
		  "auditors_url": "/v2/organizations/5eb16637-a65e-4236-9512-583317cf0387/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/5eb16637-a65e-4236-9512-583317cf0387/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/5eb16637-a65e-4236-9512-583317cf0387/domains",
		  "managers_url": "/v2/organizations/5eb16637-a65e-4236-9512-583317cf0387/managers",
		  "name": "org47",
		  "private_domains_url": "/v2/organizations/5eb16637-a65e-4236-9512-583317cf0387/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/5eb16637-a65e-4236-9512-583317cf0387/space_quota_definitions",
		  "spaces_url": "/v2/organizations/5eb16637-a65e-4236-9512-583317cf0387/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/5eb16637-a65e-4236-9512-583317cf0387/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:19Z",
		  "guid": "5eb16637-a65e-4236-9512-583317cf0387",
		  "updated_at": "2019-03-05T01:37:19Z",
		  "url": "/v2/organizations/5eb16637-a65e-4236-9512-583317cf0387"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/5d6ee56d-0a23-4c84-8d4f-9aba407efaed/app_events",
		  "auditors_url": "/v2/organizations/5d6ee56d-0a23-4c84-8d4f-9aba407efaed/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/5d6ee56d-0a23-4c84-8d4f-9aba407efaed/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/5d6ee56d-0a23-4c84-8d4f-9aba407efaed/domains",
		  "managers_url": "/v2/organizations/5d6ee56d-0a23-4c84-8d4f-9aba407efaed/managers",
		  "name": "org48",
		  "private_domains_url": "/v2/organizations/5d6ee56d-0a23-4c84-8d4f-9aba407efaed/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/5d6ee56d-0a23-4c84-8d4f-9aba407efaed/space_quota_definitions",
		  "spaces_url": "/v2/organizations/5d6ee56d-0a23-4c84-8d4f-9aba407efaed/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/5d6ee56d-0a23-4c84-8d4f-9aba407efaed/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:20Z",
		  "guid": "5d6ee56d-0a23-4c84-8d4f-9aba407efaed",
		  "updated_at": "2019-03-05T01:37:20Z",
		  "url": "/v2/organizations/5d6ee56d-0a23-4c84-8d4f-9aba407efaed"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/a61682d9-3577-4c84-8839-9947093698e8/app_events",
		  "auditors_url": "/v2/organizations/a61682d9-3577-4c84-8839-9947093698e8/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/a61682d9-3577-4c84-8839-9947093698e8/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/a61682d9-3577-4c84-8839-9947093698e8/domains",
		  "managers_url": "/v2/organizations/a61682d9-3577-4c84-8839-9947093698e8/managers",
		  "name": "org49",
		  "private_domains_url": "/v2/organizations/a61682d9-3577-4c84-8839-9947093698e8/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/a61682d9-3577-4c84-8839-9947093698e8/space_quota_definitions",
		  "spaces_url": "/v2/organizations/a61682d9-3577-4c84-8839-9947093698e8/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/a61682d9-3577-4c84-8839-9947093698e8/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:20Z",
		  "guid": "a61682d9-3577-4c84-8839-9947093698e8",
		  "updated_at": "2019-03-05T01:37:20Z",
		  "url": "/v2/organizations/a61682d9-3577-4c84-8839-9947093698e8"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/997652f9-9cd8-4e8c-9723-e8dd37e51459/app_events",
		  "auditors_url": "/v2/organizations/997652f9-9cd8-4e8c-9723-e8dd37e51459/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/997652f9-9cd8-4e8c-9723-e8dd37e51459/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/997652f9-9cd8-4e8c-9723-e8dd37e51459/domains",
		  "managers_url": "/v2/organizations/997652f9-9cd8-4e8c-9723-e8dd37e51459/managers",
		  "name": "org5",
		  "private_domains_url": "/v2/organizations/997652f9-9cd8-4e8c-9723-e8dd37e51459/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/997652f9-9cd8-4e8c-9723-e8dd37e51459/space_quota_definitions",
		  "spaces_url": "/v2/organizations/997652f9-9cd8-4e8c-9723-e8dd37e51459/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/997652f9-9cd8-4e8c-9723-e8dd37e51459/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:36:56Z",
		  "guid": "997652f9-9cd8-4e8c-9723-e8dd37e51459",
		  "updated_at": "2019-03-05T01:36:56Z",
		  "url": "/v2/organizations/997652f9-9cd8-4e8c-9723-e8dd37e51459"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/16c31361-99d0-4747-b138-a2f5641419f6/app_events",
		  "auditors_url": "/v2/organizations/16c31361-99d0-4747-b138-a2f5641419f6/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/16c31361-99d0-4747-b138-a2f5641419f6/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/16c31361-99d0-4747-b138-a2f5641419f6/domains",
		  "managers_url": "/v2/organizations/16c31361-99d0-4747-b138-a2f5641419f6/managers",
		  "name": "org50",
		  "private_domains_url": "/v2/organizations/16c31361-99d0-4747-b138-a2f5641419f6/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/16c31361-99d0-4747-b138-a2f5641419f6/space_quota_definitions",
		  "spaces_url": "/v2/organizations/16c31361-99d0-4747-b138-a2f5641419f6/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/16c31361-99d0-4747-b138-a2f5641419f6/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:37:21Z",
		  "guid": "16c31361-99d0-4747-b138-a2f5641419f6",
		  "updated_at": "2019-03-05T01:37:21Z",
		  "url": "/v2/organizations/16c31361-99d0-4747-b138-a2f5641419f6"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/70fb6e34-cde3-4958-8133-d77f757c3d69/app_events",
		  "auditors_url": "/v2/organizations/70fb6e34-cde3-4958-8133-d77f757c3d69/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/70fb6e34-cde3-4958-8133-d77f757c3d69/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/70fb6e34-cde3-4958-8133-d77f757c3d69/domains",
		  "managers_url": "/v2/organizations/70fb6e34-cde3-4958-8133-d77f757c3d69/managers",
		  "name": "org6",
		  "private_domains_url": "/v2/organizations/70fb6e34-cde3-4958-8133-d77f757c3d69/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/70fb6e34-cde3-4958-8133-d77f757c3d69/space_quota_definitions",
		  "spaces_url": "/v2/organizations/70fb6e34-cde3-4958-8133-d77f757c3d69/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/70fb6e34-cde3-4958-8133-d77f757c3d69/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:36:56Z",
		  "guid": "70fb6e34-cde3-4958-8133-d77f757c3d69",
		  "updated_at": "2019-03-05T01:36:56Z",
		  "url": "/v2/organizations/70fb6e34-cde3-4958-8133-d77f757c3d69"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/ff0517e7-37d6-41cb-8591-ff1abb2e3c90/app_events",
		  "auditors_url": "/v2/organizations/ff0517e7-37d6-41cb-8591-ff1abb2e3c90/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/ff0517e7-37d6-41cb-8591-ff1abb2e3c90/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/ff0517e7-37d6-41cb-8591-ff1abb2e3c90/domains",
		  "managers_url": "/v2/organizations/ff0517e7-37d6-41cb-8591-ff1abb2e3c90/managers",
		  "name": "org7",
		  "private_domains_url": "/v2/organizations/ff0517e7-37d6-41cb-8591-ff1abb2e3c90/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/ff0517e7-37d6-41cb-8591-ff1abb2e3c90/space_quota_definitions",
		  "spaces_url": "/v2/organizations/ff0517e7-37d6-41cb-8591-ff1abb2e3c90/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/ff0517e7-37d6-41cb-8591-ff1abb2e3c90/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:36:57Z",
		  "guid": "ff0517e7-37d6-41cb-8591-ff1abb2e3c90",
		  "updated_at": "2019-03-05T01:36:57Z",
		  "url": "/v2/organizations/ff0517e7-37d6-41cb-8591-ff1abb2e3c90"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/a3b4f0ba-ac33-499d-bea6-c32c5e1737fe/app_events",
		  "auditors_url": "/v2/organizations/a3b4f0ba-ac33-499d-bea6-c32c5e1737fe/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/a3b4f0ba-ac33-499d-bea6-c32c5e1737fe/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/a3b4f0ba-ac33-499d-bea6-c32c5e1737fe/domains",
		  "managers_url": "/v2/organizations/a3b4f0ba-ac33-499d-bea6-c32c5e1737fe/managers",
		  "name": "org8",
		  "private_domains_url": "/v2/organizations/a3b4f0ba-ac33-499d-bea6-c32c5e1737fe/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/a3b4f0ba-ac33-499d-bea6-c32c5e1737fe/space_quota_definitions",
		  "spaces_url": "/v2/organizations/a3b4f0ba-ac33-499d-bea6-c32c5e1737fe/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/a3b4f0ba-ac33-499d-bea6-c32c5e1737fe/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:36:58Z",
		  "guid": "a3b4f0ba-ac33-499d-bea6-c32c5e1737fe",
		  "updated_at": "2019-03-05T01:36:58Z",
		  "url": "/v2/organizations/a3b4f0ba-ac33-499d-bea6-c32c5e1737fe"
		}
	  },
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/d5e90699-d83b-44d3-a322-31ee72e0bc0b/app_events",
		  "auditors_url": "/v2/organizations/d5e90699-d83b-44d3-a322-31ee72e0bc0b/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/d5e90699-d83b-44d3-a322-31ee72e0bc0b/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/d5e90699-d83b-44d3-a322-31ee72e0bc0b/domains",
		  "managers_url": "/v2/organizations/d5e90699-d83b-44d3-a322-31ee72e0bc0b/managers",
		  "name": "org9",
		  "private_domains_url": "/v2/organizations/d5e90699-d83b-44d3-a322-31ee72e0bc0b/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/d5e90699-d83b-44d3-a322-31ee72e0bc0b/space_quota_definitions",
		  "spaces_url": "/v2/organizations/d5e90699-d83b-44d3-a322-31ee72e0bc0b/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/d5e90699-d83b-44d3-a322-31ee72e0bc0b/users"
		},
		"metadata": {
		  "created_at": "2019-03-05T01:36:58Z",
		  "guid": "d5e90699-d83b-44d3-a322-31ee72e0bc0b",
		  "updated_at": "2019-03-05T01:36:58Z",
		  "url": "/v2/organizations/d5e90699-d83b-44d3-a322-31ee72e0bc0b"
		}
	  }
	],
	"total_pages": 2,
	"total_results": 51
  }`

var fiftyOneOrgJSONPage2 string = `{
	"next_url": null,
	"prev_url": "/v2/organizations?order-by=name&order-direction=asc&page=1&results-per-page=50",
	"resources": [
	  {
		"entity": {
		  "app_events_url": "/v2/organizations/2afa3f15-478a-425e-ae6d-5905fc4bb4e1/app_events",
		  "auditors_url": "/v2/organizations/2afa3f15-478a-425e-ae6d-5905fc4bb4e1/auditors",
		  "billing_enabled": false,
		  "billing_managers_url": "/v2/organizations/2afa3f15-478a-425e-ae6d-5905fc4bb4e1/billing_managers",
		  "default_isolation_segment_guid": null,
		  "domains_url": "/v2/organizations/2afa3f15-478a-425e-ae6d-5905fc4bb4e1/domains",
		  "managers_url": "/v2/organizations/2afa3f15-478a-425e-ae6d-5905fc4bb4e1/managers",
		  "name": "system",
		  "private_domains_url": "/v2/organizations/2afa3f15-478a-425e-ae6d-5905fc4bb4e1/private_domains",
		  "quota_definition_guid": "8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "quota_definition_url": "/v2/quota_definitions/8b5e6b08-123e-41e3-92b9-0d9b01b98ff3",
		  "space_quota_definitions_url": "/v2/organizations/2afa3f15-478a-425e-ae6d-5905fc4bb4e1/space_quota_definitions",
		  "spaces_url": "/v2/organizations/2afa3f15-478a-425e-ae6d-5905fc4bb4e1/spaces",
		  "status": "active",
		  "users_url": "/v2/organizations/2afa3f15-478a-425e-ae6d-5905fc4bb4e1/users"
		},
		"metadata": {
		  "created_at": "2019-03-03T19:57:33Z",
		  "guid": "2afa3f15-478a-425e-ae6d-5905fc4bb4e1",
		  "updated_at": "2019-03-03T19:57:33Z",
		  "url": "/v2/organizations/2afa3f15-478a-425e-ae6d-5905fc4bb4e1"
		}
	  }
	],
	"total_pages": 2,
	"total_results": 51
  }`

var noSpaces string = `{
	"next_url": null,
	"prev_url": null,
	"resources": null,
	"total_pages": 1,
	"total_results": 0
  }`

func AddFiftyOneOrgs(server *ghttp.Server) {
	AddHandler(server, http.MethodGet, "/v2/organizations?order-by=name", http.StatusOK, []byte(fiftyOneOrgJSONPage1))
	AddHandler(server, http.MethodGet, "/v2/organizations?order-by=name&order-direction=asc&page=2&results-per-page=50", http.StatusOK, []byte(fiftyOneOrgJSONPage2))
	AddHandler(server, http.MethodGet, "/v2/spaces?order-by=name&q=organization_guid%3A6f30e06d-360e-4cd7-9849-01f28109bc37", http.StatusOK, []byte(noSpaces))
}

// /v2/organizations/6f30e06d-360e-4cd7-9849-01f28109bc37/spaces
