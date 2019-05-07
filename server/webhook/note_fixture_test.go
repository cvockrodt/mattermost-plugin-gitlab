package webhook

const MergeRequestComment = `{
	"object_kind":"note",
	"event_type":"note",
	"user":{
		"name":"manland",
		"username":"manland",
		"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
	},
	"project_id":24,
	"project":{
		"id":24,
		"name":"webhook",
		"description":"",
		"web_url":"http://localhost:3000/manland/webhook",
		"avatar_url":null,
		"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
		"git_http_url":"http://localhost:3000/manland/webhook.git",
		"namespace":"manland",
		"visibility_level":20,
		"path_with_namespace":"manland/webhook",
		"default_branch":"master",
		"ci_config_path":null,
		"homepage":"http://localhost:3000/manland/webhook",
		"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
		"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
		"http_url":"http://localhost:3000/manland/webhook.git"
	},
	"object_attributes":{
		"attachment":null,
		"author_id":50,
		"change_position":null,
		"commit_id":null,
		"created_at":"2019-04-10 10:50:44 UTC",
		"discussion_id":"5e4e2988f1bd8133cb7be80e9549a053ded5d79c",
		"id":999,
		"line_code":null,
		"note":"coucou",
		"noteable_id":37,
		"noteable_type":"MergeRequest",
		"original_position":null,
		"position":null,
		"project_id":24,
		"resolved_at":null,
		"resolved_by_id":null,
		"resolved_by_push":null,
		"st_diff":null,
		"system":false,
		"type":null,
		"updated_at":"2019-04-10 10:50:44 UTC",
		"updated_by_id":null,
		"description":"coucou",
		"url":"http://localhost:3000/manland/webhook/merge_requests/6#note_999"
	},
	"repository":{
		"name":"webhook",
		"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
		"description":"",
		"homepage":"http://localhost:3000/manland/webhook"
	},
	"merge_request":{
		"assignee_id":50,
		"author_id":1,
		"created_at":"2019-04-06 20:29:15 UTC",
		"description":"kikou",
		"head_pipeline_id":null,
		"id":37,
		"iid":6,
		"last_edited_at":null,
		"last_edited_by_id":null,
		"merge_commit_sha":null,
		"merge_error":null,
		"merge_params":{
			"force_remove_source_branch":null
		},
		"merge_status":"cannot_be_merged",
		"merge_user_id":null,
		"merge_when_pipeline_succeeds":false,
		"milestone_id":null,
		"source_branch":"master",
		"source_project_id":25,
		"state":"opened",
		"target_branch":"master",
		"target_project_id":24,
		"time_estimate":0,
		"title":"Update README.md",
		"updated_at":"2019-04-10 10:50:26 UTC",
		"updated_by_id":null,
		"url":"http://localhost:3000/manland/webhook/merge_requests/6",
		"source":{
			"id":25,
			"name":"webhook",
			"description":"",
			"web_url":"http://localhost:3000/root/webhook",
			"avatar_url":null,
			"git_ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
			"git_http_url":"http://localhost:3000/root/webhook.git",
			"namespace":"root",
			"visibility_level":20,
			"path_with_namespace":"root/webhook",
			"default_branch":"master",
			"ci_config_path":null,
			"homepage":"http://localhost:3000/root/webhook",
			"url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
			"ssh_url":"ssh://rmaneschi@localhost:2222/root/webhook.git",
			"http_url":"http://localhost:3000/root/webhook.git"
		},
		"target":{
			"id":24,
			"name":"webhook",
			"description":"",
			"web_url":"http://localhost:3000/manland/webhook",
			"avatar_url":null,
			"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
			"git_http_url":"http://localhost:3000/manland/webhook.git",
			"namespace":"manland",
			"visibility_level":20,
			"path_with_namespace":"manland/webhook",
			"default_branch":"master",
			"ci_config_path":null,
			"homepage":"http://localhost:3000/manland/webhook",
			"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
			"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
			"http_url":"http://localhost:3000/manland/webhook.git"
		},
		"last_commit":{
			"id":"a4b713084b7fcfc3570391964f6c5b5c40684a0e",
			"message":"Update README.md",
			"timestamp":"2019-04-06T20:28:44Z",
			"url":"http://localhost:3000/manland/webhook/commit/a4b713084b7fcfc3570391964f6c5b5c40684a0e",
			"author":{
				"name":"Administrator",
				"email":"admin@example.com"
			}
		},
		"work_in_progress":false,
		"total_time_spent":0,
		"human_total_time_spent":null,
		"human_time_estimate":null
	}
	}`

const IssueComment = `{
		"object_kind":"note",
		"event_type":"note",
		"user":{
			"name":"manland",
			"username":"manland",
			"avatar_url":"https://www.gravatar.com/avatar/c6b552a4cd47f7cf1701ea5b650cd2e3?s=80\\u0026d=identicon"
		},
		"project_id":24,
		"project":{
			"id":24,
			"name":"webhook",
			"description":"",
			"web_url":"http://localhost:3000/manland/webhook",
			"avatar_url":null,
			"git_ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
			"git_http_url":"http://localhost:3000/manland/webhook.git",
			"namespace":"manland",
			"visibility_level":20,
			"path_with_namespace":"manland/webhook",
			"default_branch":"master",
			"ci_config_path":null,
			"homepage":"http://localhost:3000/manland/webhook",
			"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
			"ssh_url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
			"http_url":"http://localhost:3000/manland/webhook.git"
		},
		"object_attributes":{
			"attachment":null,
			"author_id":50,
			"change_position":null,
			"commit_id":null,
			"created_at":"2019-04-10 10:44:13 UTC",
			"discussion_id":"f86a43ac361b0d51b3a1c12b1c178ab060097a2d",
			"id":997,
			"line_code":null,
			"note":"coucou3",
			"noteable_id":181,
			"noteable_type":"Issue",
			"original_position":null,
			"position":null,
			"project_id":24,
			"resolved_at":null,
			"resolved_by_id":null,
			"resolved_by_push":null,
			"st_diff":null,
			"system":false,
			"type":null,
			"updated_at":"2019-04-10 10:44:13 UTC",
			"updated_by_id":null,
			"description":"coucou3",
			"url":"http://localhost:3000/manland/webhook/issues/1#note_997"
		},
		"repository":{
			"name":"webhook",
			"url":"ssh://rmaneschi@localhost:2222/manland/webhook.git",
			"description":"",
			"homepage":"http://localhost:3000/manland/webhook"
		},
		"issue":{
			"author_id":1,
			"closed_at":null,
			"confidential":false,
			"created_at":"2019-04-06 21:03:04 UTC",
			"description":"hello world!",
			"due_date":null,
			"id":181,
			"iid":1,
			"last_edited_at":null,
			"last_edited_by_id":null,
			"milestone_id":null,
			"moved_to_id":null,
			"project_id":24,
			"relative_position":1073742323,
			"state":"opened",
			"time_estimate":0,
			"title":"test new issue",
			"updated_at":"2019-04-10 10:44:13 UTC",
			"updated_by_id":null,
			"url":"http://localhost:3000/manland/webhook/issues/1",
			"total_time_spent":0,
			"human_total_time_spent":null,
			"human_time_estimate":null,
			"assignee_ids":[50],
			"assignee_id":50
		}
		}`