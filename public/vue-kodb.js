Vue.component("kodb", {
        data: function() {
                return {
                        selectedLibrary: null,
                        librarys:[
                        ],
                        librarisData:{
                        },
                }
        },
        webSockets: {
                connected() {
                        this.$wsocket.send({
                                "command": "getSchema"
                        })
                },
                command: {
                        setSchema(msg) {
                                for (let l of msg.librarys) {
                                        if (null == this.librarisData[l.name]) {
                                                Vue.set(this.librarisData, l.name, [])

                                                this.$wsocket.send({
                                                        "command": "getLibraryRows",
                                                        "library": l.name
                                                })
                                        }
                                }

                                this.librarys = msg.librarys
                        },
                        setLibraryRows(msg) {
                                const rows = this.librarisData[msg.library]
                                if (null == rows) {
                                        return
                                }
                                rows.splice(0, rows.length)
                                rows.push(...msg.rows)
                        },
                        newRow(msg) {
                                const rows = this.librarisData[msg.library]
                                if (null == rows) {
                                        return
                                }
                                rows.push({
                                        "rowId": msg.rowId,
                                        "data": {}
                                })
                        },
                        deleteRow(msg) {
                                const rows = this.librarisData[msg.library]
                                if (null == rows) {
                                        return
                                }
                                const rowId = msg.rowId
                                const rowIndex = rows.findIndex(x => x.rowId == rowId)
                                if (-1 != rowIndex) {
                                        rows.splice(rowIndex, 1)
                                }
                        },
                        updateValue(msg) {
                                const rows = this.librarisData[msg.library]
                                if (null == rows) {
                                        return
                                }
                                const rowId = msg.rowId
                                const columnId = msg.columnId
                                const rowIndex = rows.findIndex(x => x.rowId == rowId)
                                if (-1 == rowIndex) {
                                        return;
                                }
                                const row = rows[rowIndex]
                                const columnData = row.data[columnId]
                                
                                Vue.set(columnData, "exists", msg.exists)
                                Vue.set(columnData, "value", msg.value)
                        }
                }
        },
        template:
`
<v-app>
        <v-app-bar
                app
                clipped-left
                dense  
        >
                <v-app-bar-nav-icon></v-app-bar-nav-icon>

                <v-toolbar-title>Page title</v-toolbar-title>

                <v-spacer></v-spacer>

                <v-btn icon>
                        <v-icon>mdi-magnify</v-icon>
                </v-btn>

                <v-btn icon>
                        <v-icon>mdi-dots-vertical</v-icon>
                </v-btn>
        

                <template v-slot:extension>
                        <v-tabs
                                v-model="selectedLibrary"
                                show-arrows
                        >
                                <v-tab
                                        v-for="l in librarys"
                                        :key="l.id"
                                >
                                        {{ l.name }}
                                </v-tab>
                        </v-tabs>
                </template>
        </v-app-bar>

        <v-content>
                <v-container>
                        <v-tabs-items
                                v-model="selectedLibrary"
                        >
                                <v-tab-item
                                        v-for="l in librarys"
                                        :key="l.id"
                                >
                                        <kodb-library
                                                v-bind:librarySchema="l"
                                                v-bind:rows="librarisData[l.name]"
                                        >
                                        </kodb-library>
                                </v-tab-item>
                        </v-tabs-items>
                </v-container>
        </v-content>
</v-app>
`
})