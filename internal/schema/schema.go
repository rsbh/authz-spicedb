package schema

const schema = `
definition blog/user {}

definition blog/post {
    relation reader: blog/user
    relation writer: blog/user

    permission read = reader + writer
    permission write = writer
}
`