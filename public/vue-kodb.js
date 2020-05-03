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

                getLibraryColumns(libraryName, showHidden) {
                        const library = this.schema.map[libraryName]
                        if (null == library) {
                                console.warn(`library <${libraryName}> not exists`)
                                return {}
                        }
                        if (showHidden) {
                                return library.columns
                        }
                        return library.columns.filter(c => !c.hidden)
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

                getColumnName(libraryName, columnId) {
                        return this.getColumnData(libraryName, columnId).name
                },

                getColumnType(libraryName, columnId) {
                        return this.getColumnData(libraryName, columnId).type
                },

                isParentOf(libraryName, rowId, columnId, rowData) {
                        const data = rowData.data
                        
                        const parentLibrary = data['parentLibrary']
                        if (null == parentLibrary
                                || libraryName != parentLibrary.value
                        ) return false

                        const parentRow = data['parentRow']
                        if (null == parentRow
                                || rowId != parentRow.value
                        ) return false

                        const parentColumn = data['parentColumn']
                        if (null == parentColumn
                                || columnId != parentColumn.value
                        ) return
                        
                        return true
                }
        },
})

Vue.component("kodb", {
        data: function() {
                return {
                        selectedLibrary: null,
                        schema: {
                                types: {},
                                list: [],
                                map: {},
                                rowsMap: {}
                        },
                }
        },
        webSockets: {
                connected() {
                        this.$wsocket.send({
                                "command": "getTypes"
                        })
                        this.$wsocket.send({
                                "command": "getSchema"
                        })
                },
                command: {
                        setTypes(msg) {
                                this.schema.types = msg.types
                        },
                        setSchema(msg) {
                                
                                
                                const newList = []
                                const newMap = {}
                                
                                for (let l of msg.librarys) {
                                        // HACK!!!
                                        l.value = l.id

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
                                const rows = this.schema.rowsMap[msg.library]
                                if (null == rows) {
                                        return
                                }
                                const library = this.schema.map[msg.library]
                                if (null == library) {
                                        return
                                }

                                rows.push({
                                        "rowId": msg.rowId
                                })
                        },
                        deleteRow(msg) {
                                const rows = this.schema.rowsMap[msg.library]
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
                                const rows = this.schema.rowsMap[msg.library]
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
                                if (null == row.data[columnId]) {
                                        Vue.set(row.data, columnId, {})
                                }

                                const columnData = row.data[columnId]
                                
                                Vue.set(columnData, "exists", msg.exists)
                                Vue.set(columnData, "value", msg.value)
                                Vue.set(columnData, "error", msg.error)
                        },
                        swapRows(msg) {
                                const rows = this.schema.rowsMap[msg.library]
                                if (null == rows) {
                                        return
                                }
                                const row = rows[msg.j]
                                const row0 = rows[msg.i]
                                if (row.rowId == msg.row && row0.rowId == msg.row0) {
                                        Vue.set(rows, msg.i, row)
                                        Vue.set(rows, msg.j, row0)
                                } else {
                                        log.warn("Invalidate library", msg.library)

                                        this.$wsocket.send({
                                                "command": "getLibraryRows",
                                                "library": msg.library
                                        })
                                }
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
                
                <kodb-types-dialog
                        :types="schema.types"
                >
                </kodb-types-dialog>            

                <template v-slot:extension>
                        <v-tabs dense
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

        <v-content >
                <v-container fluid>
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