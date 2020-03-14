Vue.component("vue-kodb-schema-literal-column", {
        props: [
        ],
        template:
`
<v-card>
        <v-col>
                <v-text-field
                >
                </v-text-field>    

                <v-radio-group>
                <v-radio
                        label="String"
                ></v-radio>
                <v-radio
                        label="Text"
                ></v-radio>
                <v-radio
                        label="Int"
                ></v-radio>
                <v-radio
                        label="Float"
                ></v-radio>
                <v-radio
                        label="Option"
                ></v-radio>
                <v-radio
                        label="Set"
                ></v-radio>
                </v-radio-group>
        </v-col>
</v-card>
`})

Vue.component("vue-kodb-schema-ref-column", {
        props: [
                "schema"
        ],
        template:
`
<v-card>
        <v-col>
                <v-text-field
                >
                </v-text-field>

                <v-select
                        :items="schema"
                >
                </v-select>
        </v-col>
</v-card>
`})

Vue.component("vue-kodb-schema-list-column", {
        props: [
                "schema"
        ],
        template:
`
<v-card>
        <v-col>
                <v-text-field
                >
                </v-text-field>

                <v-select
                        :items="schema"
                >
                </v-select>
        </v-col>
</v-card>
`})