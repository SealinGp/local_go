package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

/*
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.12.md
数据加密
*/

var encryptFuncs = map[string]func(){
	"enc1":  enc1,
	"Md5En": Md5En,
}

var body string = `{
  "ref": "refs/heads/master",
  "before": "a28418b92902ee9dfc111ef3168daab2210239cb",
  "after": "f2706a4e4a6f57906abce5a2df96988847328156",
  "repository": {
    "id": 207484708,
    "node_id": "MDEwOlJlcG9zaXRvcnkyMDc0ODQ3MDg=",
    "name": "billing",
    "full_name": "SealinGp/billing",
    "private": true,
    "owner": {
      "name": "SealinGp",
      "email": "32982627+SealinGp@users.noreply.github.com",
      "login": "SealinGp",
      "id": 32982627,
      "node_id": "MDQ6VXNlcjMyOTgyNjI3",
      "avatar_url": "https://avatars3.githubusercontent.com/u/32982627?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/SealinGp",
      "html_url": "https://github.com/SealinGp",
      "followers_url": "https://api.github.com/users/SealinGp/followers",
      "following_url": "https://api.github.com/users/SealinGp/following{/other_user}",
      "gists_url": "https://api.github.com/users/SealinGp/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/SealinGp/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/SealinGp/subscriptions",
      "organizations_url": "https://api.github.com/users/SealinGp/orgs",
      "repos_url": "https://api.github.com/users/SealinGp/repos",
      "events_url": "https://api.github.com/users/SealinGp/events{/privacy}",
      "received_events_url": "https://api.github.com/users/SealinGp/received_events",
      "type": "User",
      "site_admin": false
    },
    "html_url": "https://github.com/SealinGp/billing",
    "description": "Billing System",
    "fork": false,
    "url": "https://github.com/SealinGp/billing",
    "forks_url": "https://api.github.com/repos/SealinGp/billing/forks",
    "keys_url": "https://api.github.com/repos/SealinGp/billing/keys{/key_id}",
    "collaborators_url": "https://api.github.com/repos/SealinGp/billing/collaborators{/collaborator}",
    "teams_url": "https://api.github.com/repos/SealinGp/billing/teams",
    "hooks_url": "https://api.github.com/repos/SealinGp/billing/hooks",
    "issue_events_url": "https://api.github.com/repos/SealinGp/billing/issues/events{/number}",
    "events_url": "https://api.github.com/repos/SealinGp/billing/events",
    "assignees_url": "https://api.github.com/repos/SealinGp/billing/assignees{/user}",
    "branches_url": "https://api.github.com/repos/SealinGp/billing/branches{/branch}",
    "tags_url": "https://api.github.com/repos/SealinGp/billing/tags",
    "blobs_url": "https://api.github.com/repos/SealinGp/billing/git/blobs{/sha}",
    "git_tags_url": "https://api.github.com/repos/SealinGp/billing/git/tags{/sha}",
    "git_refs_url": "https://api.github.com/repos/SealinGp/billing/git/refs{/sha}",
    "trees_url": "https://api.github.com/repos/SealinGp/billing/git/trees{/sha}",
    "statuses_url": "https://api.github.com/repos/SealinGp/billing/statuses/{sha}",
    "languages_url": "https://api.github.com/repos/SealinGp/billing/languages",
    "stargazers_url": "https://api.github.com/repos/SealinGp/billing/stargazers",
    "contributors_url": "https://api.github.com/repos/SealinGp/billing/contributors",
    "subscribers_url": "https://api.github.com/repos/SealinGp/billing/subscribers",
    "subscription_url": "https://api.github.com/repos/SealinGp/billing/subscription",
    "commits_url": "https://api.github.com/repos/SealinGp/billing/commits{/sha}",
    "git_commits_url": "https://api.github.com/repos/SealinGp/billing/git/commits{/sha}",
    "comments_url": "https://api.github.com/repos/SealinGp/billing/comments{/number}",
    "issue_comment_url": "https://api.github.com/repos/SealinGp/billing/issues/comments{/number}",
    "contents_url": "https://api.github.com/repos/SealinGp/billing/contents/{+path}",
    "compare_url": "https://api.github.com/repos/SealinGp/billing/compare/{base}...{head}",
    "merges_url": "https://api.github.com/repos/SealinGp/billing/merges",
    "archive_url": "https://api.github.com/repos/SealinGp/billing/{archive_format}{/ref}",
    "downloads_url": "https://api.github.com/repos/SealinGp/billing/downloads",
    "issues_url": "https://api.github.com/repos/SealinGp/billing/issues{/number}",
    "pulls_url": "https://api.github.com/repos/SealinGp/billing/pulls{/number}",
    "milestones_url": "https://api.github.com/repos/SealinGp/billing/milestones{/number}",
    "notifications_url": "https://api.github.com/repos/SealinGp/billing/notifications{?since,all,participating}",
    "labels_url": "https://api.github.com/repos/SealinGp/billing/labels{/name}",
    "releases_url": "https://api.github.com/repos/SealinGp/billing/releases{/id}",
    "deployments_url": "https://api.github.com/repos/SealinGp/billing/deployments",
    "created_at": 1568097947,
    "updated_at": "2019-12-09T14:17:00Z",
    "pushed_at": 1575902023,
    "git_url": "git://github.com/SealinGp/billing.git",
    "ssh_url": "git@github.com:SealinGp/billing.git",
    "clone_url": "https://github.com/SealinGp/billing.git",
    "svn_url": "https://github.com/SealinGp/billing",
    "homepage": null,
    "size": 85550,
    "stargazers_count": 0,
    "watchers_count": 0,
    "language": "Vue",
    "has_issues": true,
    "has_projects": true,
    "has_downloads": true,
    "has_wiki": true,
    "has_pages": false,
    "forks_count": 1,
    "mirror_url": null,
    "archived": false,
    "disabled": false,
    "open_issues_count": 0,
    "license": null,
    "forks": 1,
    "open_issues": 0,
    "watchers": 0,
    "default_branch": "master",
    "stargazers": 0,
    "master_branch": "master"
  },
  "pusher": {
    "name": "SealinGp",
    "email": "32982627+SealinGp@users.noreply.github.com"
  },
  "sender": {
    "login": "SealinGp",
    "id": 32982627,
    "node_id": "MDQ6VXNlcjMyOTgyNjI3",
    "avatar_url": "https://avatars3.githubusercontent.com/u/32982627?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/SealinGp",
    "html_url": "https://github.com/SealinGp",
    "followers_url": "https://api.github.com/users/SealinGp/followers",
    "following_url": "https://api.github.com/users/SealinGp/following{/other_user}",
    "gists_url": "https://api.github.com/users/SealinGp/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/SealinGp/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/SealinGp/subscriptions",
    "organizations_url": "https://api.github.com/users/SealinGp/orgs",
    "repos_url": "https://api.github.com/users/SealinGp/repos",
    "events_url": "https://api.github.com/users/SealinGp/events{/privacy}",
    "received_events_url": "https://api.github.com/users/SealinGp/received_events",
    "type": "User",
    "site_admin": false
  },
  "created": false,
  "deleted": false,
  "forced": false,
  "base_ref": null,
  "compare": "https://github.com/SealinGp/billing/compare/a28418b92902...f2706a4e4a6f",
  "commits": [
    {
      "id": "f2706a4e4a6f57906abce5a2df96988847328156",
      "tree_id": "97fb2b9372b4abb803c8642f73ec1ac3d2bac8b5",
      "distinct": true,
      "message": "[update]webhook test",
      "timestamp": "2019-12-09T22:33:20+08:00",
      "url": "https://github.com/SealinGp/billing/commit/f2706a4e4a6f57906abce5a2df96988847328156",
      "author": {
        "name": "sealingp",
        "email": "sealingp@163.com",
        "username": "SealinGp"
      },
      "committer": {
        "name": "sealingp",
        "email": "sealingp@163.com",
        "username": "SealinGp"
      },
      "added": [

      ],
      "removed": [

      ],
      "modified": [
        "backend/billing.go"
      ]
    }
  ],
  "head_commit": {
    "id": "f2706a4e4a6f57906abce5a2df96988847328156",
    "tree_id": "97fb2b9372b4abb803c8642f73ec1ac3d2bac8b5",
    "distinct": true,
    "message": "[update]webhook test",
    "timestamp": "2019-12-09T22:33:20+08:00",
    "url": "https://github.com/SealinGp/billing/commit/f2706a4e4a6f57906abce5a2df96988847328156",
    "author": {
      "name": "sealingp",
      "email": "sealingp@163.com",
      "username": "SealinGp"
    },
    "committer": {
      "name": "sealingp",
      "email": "sealingp@163.com",
      "username": "SealinGp"
    },
    "added": [

    ],
    "removed": [

    ],
    "modified": [
      "backend/billing.go"
    ]
  }
}`

//sha1
func enc1() {
	//github
	//69d7d753e0ca7c448afa10f10a8c09d8638ce38d
	SHA1 := sha1.New()
	Msg := "billingPush"
	MsgBy := []byte(Msg)
	Key := "billingPush"
	KeyBy := []byte(Key)

	//加密
	_, err := SHA1.Write(MsgBy)
	if err != nil {
		fmt.Println(err)
		return
	}
	//checksum    := SHA1.Sum(nil)
	//checksumStr := fmt.Sprintf("%x",checksum)
	//fmt.Println(checksumStr)

	//hmac的sha1
	HmacWithSha1 := hmac.New(sha1.New, KeyBy)
	HmacWithSha1.Write([]byte(body))
	checksum1 := HmacWithSha1.Sum(nil)
	checksum1Str := hex.EncodeToString(checksum1)
	fmt.Println(checksum1Str)

	//a   := "69d7d753e0ca7c448afa10f10a8c09d8638ce38d"
	//b,_ := hex.DecodeString(a)
	//fmt.Println(hmac.Equal(b,[]byte(checksum1)))
}

/**
https://studygolang.com/articles/2283
md5加密
*/
func Md5En() {
	msg := "121"
	m := md5.New()
	m.Write([]byte(msg))
	by := m.Sum(nil)
	fmt.Println(hex.EncodeToString(by))
}
