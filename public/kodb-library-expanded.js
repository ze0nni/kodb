Vue.component("kodb-library-expanded", {
    props: [
        "depth",
        "schema",
        "libraryName",
        
        "parentLibraryName",
        "parentColumnId",
        "parentRowId",
    ],
    data() {
        return {
                selectedRow: [],

                expandedLibraryName: null,
                expandedColumnId: null,
        }
    },
    computed: {
        columns() {
                return this.getLibraryColumns(this.libraryName)
        },
        rows() {
                const rows = this.schema.rowsMap[this.libraryName]
                const parentLibraryId = this.parentLibraryName
                const parentRowId = this.parentRowId
                const parentColumnId = this.parentColumnId
                return (rows||[])
                        .filter(r => this.isParentOf(
                                parentLibraryId,
                                parentRowId,
                                parentColumnId,
                                r
                        ))
        },
        expandedRows() {
                const columns = this.columns
                return this.rows
                        .map(r => [
                                {extendedRow: false, row: r, columns: columns},
                                { extendedRow: true, row: r, columns:[], numColumns: columns.length }
                        ])
                        .reduce((a, b) => a.concat(b), [])
        }
    },
    methods: {
            isSelected(row) {
                    return -1 != this.selectedRow.indexOf(row)
            },
            toggleSelect(row) {
                const selected = this.selectedRow
                if (0 == selected.length) {
                        selected.push(row)
                } else if (row == selected[0]) {
                        selected.splice(0, selected.length)
                } else {
                        selected.splice(0, selected.length, row)
                }
            },
            expandRow(library, columnId) {
                    return () => {
                        if (this.expandedLibraryName == library) {
                                this.expandedLibraryName = null
                        } else {
                                this.expandedLibraryName = library
                                this.expandedColumnId = columnId
                        }
                    }
            },
            colorFromDepth(depth) {
                    return 'grey lighten-2'
            }
    },
    template:
`
<v-card dense tile :color=colorFromDepth(depth)>
        <v-simple-table dense>
                <thead>
                        <tr color="red">
                                <th :colspan="columns.length + 1"
                                        style="margin:0; padding:0"
                                >
                                        <v-toolbar tile dark dense>
                                                <v-toolbar-title>{{ getColumnName(parentLibraryName, parentColumnId) }}</v-toolbar-title>

                                                <v-spacer></v-spacer>
                                                
                                                <kodb-library-rows-menu
                                                        :libraryName="libraryName"
                                                        :libraryRows="rows"

                                                        :parentLibraryName="parentLibraryName"
                                                        :parentRowId="parentRowId"
                                                        :parentColumnId="parentColumnId"

                                                        :selectedRows="selectedRow"
                                                >
                                                </kodb-library-rows-menu>
                                        </v-toolbar>
                                </th>
                        </tr>
                        <tr>
                                <th style="width:1em; min-wdth:1em">
                                        <v-icon>
                                                mdi-chevron-up
                                        </v-icon>
                                </th>
                                <th v-for="col in columns"
                                >
                                        {{ col.name }}
                                </th>
                        </tr>
                </thead>
                <tbody>
                        <tr v-for="r in expandedRows"
                        >
                                <td v-if="!r.extendedRow">
                                        <v-icon v-on:click="toggleSelect(r.row)">
                                                {{ isSelected(r.row) ? "mdi-check-box-outline" : "mdi-checkbox-blank-outline" }}
                                        </v-icon>
                                </td>
                                <td v-for="col in r.columns"
                                >
                                        <kodb-library-cell

                                                :schema="schema"
                                                :libraryName="libraryName"
                                                :rowId="r.row.rowId"
                                                :columnId="col.id"

                                                :rowData="r.row"
                                                :cellData="r.row.data[col.value]"

                                                :expandRow="expandRow(col.reference, col.id)"
                                                :isExpanded="false"
                                        >
                                        </kodb-library-cell>
                                </td>

                                <td v-if="r.extendedRow && r.extendedRow && expandedLibraryName"></td>
                                <td v-if="r.extendedRow && expandedLibraryName"
                                        :colspan="r.numColumns">
                                        <kodb-library-expanded
                                                v-if="expandedLibraryName"

                                                :depth="Number(depth)+1"
                                                :schema="schema"
                                                :libraryName="expandedLibraryName"
                                                
                                                :parentLibraryName="libraryName"
                                                :parentColumnId="expandedColumnId"
                                                :parentRowId="r.row.rowId"
                                        >
                                        </kodb-library-expanded>
                                </td>
                        </tr>
                </tbody>
        </v-simple-table>
</v-card>
`
});