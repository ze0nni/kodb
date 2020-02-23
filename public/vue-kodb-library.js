Vue.component("kodb-library", {
        props: [
                "librarySchema"
        ],
        data: function() {
                return {
                        rows:[],
                }
        },
        webSockets: {
                connected() {
                        this.$wsocket.send({
                                "command": "getLibraryRows",
                                "library": this.librarySchema.name
                        })
                },
                command: {
                        setLibraryRows(msg) {
                                if (msg.library == this.librarySchema.name) {
                                        this.rows = msg.rows
                                }
                        }
                }
        },
        template:
`
<v-data-table
        :headers="librarySchema.columns"
        :items="rows"
        :items-per-page="10"
></v-data-table>
`
});