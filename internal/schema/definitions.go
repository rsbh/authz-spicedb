package schema

const schema = `
definition user {}

definition group {
   relation member : user 
   permission view =  member
}

definition resource/firehose {
   relation manager: group | project#firehose_admins
   permission manage = manager + manager->member + manager->firehose_admins
}

definition resource/beast {
   relation manager: group | project#beast_admins
   permission manage = manager + manager->member + manager->beast_admins
}

definition project {
   relation manager: user
   relation firehose_admins: group#member
   relation beast_admins: group#member
   permission manage = manager
}

definition organization {
   relation manager: user
   permission manage = manager
}
`
