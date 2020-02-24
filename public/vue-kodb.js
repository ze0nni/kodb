Vue.component("kodb", {
        data: function() {
                return {
                        selectedLibrary: null,
                        librarys:[
                        ]
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
                                this.librarys = msg.librarys
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
                                        />
                                </v-tab-item>
                        </v-tabs-items>
                </v-container>
        </v-content>
</v-app>
`
})