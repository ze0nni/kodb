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
                newField() {
                        this.$wsocket.send({
                                "command": "newField",
                                "type": this.type.name
                        })
                }
        },
        template:
`
<v-simple-table>
        <tr>
                <td>
                        Default
                </td>
                <td>
                        <v-chip-group column>
                                <kodb-type-field-view v-for="f in type.fields"
                                        :key="f.id"

                                        :type="type"
                                        :field="f"
                                >
                                </kodb-type-field-view>
                                <v-chip color="green" @click="newField">
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
                <v-list-item @click="">
                        <v-list-item-title>Move left</v-list-item-title>
                </v-list-item>
                <v-list-item @click="">
                        <v-list-item-title>Move right</v-list-item-title>
                </v-list-item>

                <v-list-item @click="">
                        <v-list-item-title>Delete</v-list-item-title>
                </v-list-item>
        </v-list>
</v-menu>
`
})