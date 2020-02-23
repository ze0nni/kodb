Vue.component("kodb", {
        data: function() {
                return {
                        selectedLibrary: null,
                        librarys:[
                                {id: 1, name:"Users"},
                                {id: 2, name:"Inventory"},
                        ]
                }
        },
        template:
`
<v-app>
        <v-content>
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

                        <v-tab-item
                                v-for="l in librarys"
                                :key="l.id"
                        >
                                <kodb-library/>
                        </v-tab-item>
                </v-tabs>
        </v-content>
</v-app>
`
})