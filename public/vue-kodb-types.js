Vue.component("kodb-types-dialog", {
        props: [
                "types"
        ],
        computed: {
                names() {
                        return Object.keys(this.types)
                }
        },
        template:
`
<v-dialog>
        <template v-slot:activator="{ on }">
        <v-btn text v-on="on">
                <v-icon left>mdi-shape-outline</v-icon>
                Types
        </v-btn>
        </template>

        <v-card dense>
                <v-toolbar dark>
                        Types
                        <v-spacer></v-spacer>
                        <v-btn text outlined>
                                New type
                        </v-btn>
                </v-toolbar>
                <v-tabs dense>
                        <v-tab v-for="n in names"
                                :key="n"

                                dense
                        >
                                {{ n }}
                        </v-tab>

                        <v-tab-item v-for="n in names"
                                :key="n"

                                dense
                        >
                                <kodb-type-view
                                        :type="types[n]"
                                >
                                </kodb-type-view>
                        </v-tab-item>
                </v-tabs>
        </v-card>
</v-dialog>
`
})

Vue.component("kodb-type-view", {
        props:[
                "type"
        ],
        methods: {
                fieldsForCase(c) {
                        return Object
                                .values(this.type.fields)
                                .filter(f => f['case'] == c)
                },

                newField(fieldCase) {
                        this.$wsocket.send({
                                "command": "newField",
                                "type": this.type.name,
                                "case": fieldCase
                        })
                }
        },
        template:
`
<v-simple-table>
        <tr v-for="c in type.cases"
                :key="c"
        >
                <td class="text-right">
                        {{ c || 'Default' }}
                </td>
                <td>
                        <v-chip-group column>
                                <kodb-type-field-view v-for="f in fieldsForCase(c)"
                                        :key="f.id"

                                        :type="type"
                                        :field="f"
                                >
                                </kodb-type-field-view>
                                <v-chip color="green" @click="newField(c)">
                                        +
                                </v-chip>
                        </v-chip-group>
                </td>
        </tr>
        <tr>
                <td>
                        <v-btn block outlined text>
                                New case
                        </v-btn>
                </td>
                <td>
                </td>
        </tr>
</v-simple-table>
`
})

Vue.component("kodb-type-field-view", {
        props:[
                "type",
                "field"
        ],
        methods: {
                deleteField() {
                        this.$wsocket.send({
                                "command": "deleteField",
                                "type": this.type.name,
                                "field": this.field.id,
                        })
                }
        },
        template:
`
<v-menu offset-y>
        <template v-slot:activator="{ on }">
                <v-chip v-on="on">
                        {{field.name}}
                </v-chip>
        </template>

        <v-list>
                <kodb-type-field-view-rename
                        :field="field"
                        :type="type"
                >
                </kodb-type-field-view-rename>

                <v-list-item @click="">
                        <v-list-item-title>Move left</v-list-item-title>
                </v-list-item>
                <v-list-item @click="">
                        <v-list-item-title>Move right</v-list-item-title>
                </v-list-item>

                <v-list-item @click="deleteField">
                        <v-list-item-title>Delete</v-list-item-title>
                </v-list-item>
        </v-list>
</v-menu>
`
})

Vue.component("kodb-type-field-view-rename", {
        props:[
                "type",
                "field"
        ],
        data() {
                return {
                        dialog: false,

                        dialogName: "",
                }
        },
        methods: {
                save() {
                        this.dialog = false

                        this.$wsocket.send({
                                "command": "updateField",

                                "type": this.type.name,
                                "field": this.field.id,

                                "name": this.dialogName
                        })
                }
        },
        watch: {
                dialog(value) {
                        if (value) {
                                this.dialogName = this.field.name
                        }
                }
        },
        template:
`
<v-dialog v-model="dialog">
        <template v-slot:activator="{ on }">
                <v-list-item v-on="on">Rename</v-list-item>
        </template>
        <v-card>
                <v-toolbar dark>
                        Edit "{{type.name}}.{{field.name}}"
                </v-toolbar>
                <v-card-text>
                        <v-text-field
                                v-model="dialogName"
                                label="name"
                        ></v-text-field>
                </v-card-text>
                <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn text @click="save">
                                Save
                        </v-btn>
                </v-card-actions>
        </v-card>
</v-dialog>
`
})