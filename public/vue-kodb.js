Vue.mixin({
        methods: {
                getRowData(row, col) {
                        const data = row.data
                        if (null == data) return undefined
                        return data[col]
                },

                getRowValue(row, col) {
                        const data = this.getRowData(row, col)
                        if (null == data) return undefined
                        if (data.exists) {
                                return data.value
                        }

                        return undefined
                },

                getColumnsOf(libraryName) {
                        console.log(schema)
                        return []
                },

                getColumnData(libraryName, columnId) {
                        const library = this.schema.map[libraryName]
                        if (null == library) {
                                console.warn(`library <${libraryName}> not exists`)
                                return {}
                        }
                        const column = library.columnsMap[columnId]
                        if (null == column) {
                                console.warn(`column <${columnId}>library <${libraryName}> not exists`)
                                return {}
                        }
                        return column
                },

                getColumnType(libraryName, columnId) {
                        return this.getColumnData(libraryName, columnId).type
                },
        },
})

Vue.component("kodb", {
        data: function() {
                return {
                        selectedLibrary: null,
                        schema: {
                                list: [],
                                map: {},
                                rowsMap: {}
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
                                
                                
                                const newList = []
                                const newMap = {}
                                
                                for (let l of msg.librarys) {
                                        // HACK!!!
                                        l.value = l.id

                                        // HACK
                                        l.columnsMap = {

                                        }

                                        for (let c of l.columns) {
                                                c.value = c.id
                                                l.columnsMap[c.id] = c
                                        }

                                        newList.push(l)
                                        newMap[l.name] = l
                                                
                                        if (undefined == this.schema.rowsMap[l.name]) {
                                                Vue.set(this.schema.rowsMap, l.name, [])
                                        }

                                        this.$wsocket.send({
                                                "command": "getLibraryRows",
                                                "library": l.name
                                        })
                                }

                                this.schema.list = newList
                                this.schema.map = newMap

                        },
                        setLibraryRows(msg) {
                                const rows = this.schema.rowsMap[msg.library]
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
                                Vue.set(columnData, "error", msg.error)
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

                <v-menu offset-y left>
                        <template v-slot:activator="{ on }">
                                <v-btn icon v-on="on">
                                        <v-icon>mdi-dots-vertical</v-icon>
                                </v-btn>
                        </template>

                        <v-list>
                                <v-list-item>
                                        <kodb-schema-manager
                                                :schema="schema"
                                        >
                                        </kodb-schema-manager>
                                </v-list-item>
                      </v-list>
                </v-menu>
        

                <template v-slot:extension>
                        <v-tabs
                                v-model="selectedLibrary"
                                show-arrows
                        >
                                <v-tab
                                        v-for="l in schema.list"
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
                                        v-for="l in schema.list"
                                        :key="l.id"
                                >
                                        <kodb-library
                                                v-bind:schema="schema"
                                                v-bind:libraryName="l.name"
                                        >
                                        </kodb-library>
                                </v-tab-item>
                        </v-tabs-items>
                </v-container>
        </v-content>
</v-app>
`
})