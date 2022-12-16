<template>
  <b-tab title="TabConfigurator">
    <div
      v-if="alert"
      class="custom-warning text-h3"
    >
      <p> {{ msg }} </p>
    </div>

    <div
      class="d-flex align-items-center"
    >
      <h3
        class="font-weight-light text-primary"
      >
        {{ $t('tab.title') }}
      </h3>

      <b-button
        variant="link"
        size="lg"
        :title="$t('tab.tooltip.addTab')"
        class="h3"
        @click="Add"
      >
        {{ $t('tab.addTab') }}
      </b-button>
    </div>

    <div v-if="block.options.tabs.length">
      <draggable
        v-model="block.options.tabs"
        :disabled="editFocused"
      >
        <transition-group tag="div">
          <b-row
            v-for="(tab, index) in block.options.tabs"
            :key="index"
            class=" d-flex justify-content-around mb-3 mt-3 ml-0 mr-0"
          >
            <b-col
              cols="1"
              class="mt-2"
            >
              <b-button
                variant="link"
                :title="$t('tab.tooltip.move')"
              >
                <font-awesome-icon
                  :icon="['fas', 'bars']"
                  class="text-secondary grab h4"
                />
              </b-button>
            </b-col>

            <b-col
              cols="7"
              class=" p-0 text-center d-flex justify-content-center align-items-center"
            >
              <h4
                v-if="tabMode[index] === 'view'"
              >
                {{ tab.block.title }}
              </h4>

              <b-col
                v-else
                cols="6"
                class="p-0"
              >
                <b-form-group>
                  <b-form-select
                    v-model="selectedBlocks[index]"
                    :options="options"
                    size="sm"
                    @change="createTab(index)"
                  >
                    <template #first>
                      <b-form-select-option
                        :value="null"
                        disabled
                      >
                        {{ $t('tab.selectBlock') }}
                      </b-form-select-option>
                    </template>
                  </b-form-select>
                </b-form-group>
                <div
                  v-if="tabWarning"
                  class="custom-warning "
                >
                  <p>{{ msg }}</p>
                </div>
              </b-col>

              <b-button
                v-if="tabMode[index] === 'view'"
                :title="$t('tab.tooltip.changeTab')"
                variant="link"
                class="mb-2"
                @click="switchTabMode(index, 'choose')"
              >
                <font-awesome-icon
                  :icon="['far', 'edit']"
                />
              </b-button>

              <b-button
                v-if="tabMode[index] === 'choose'"
                :title="$t('tab.tooltip.cancel')"
                variant="link"
                class="mb-3"
                @click="switchTabMode(index, 'view')"
              >
                <font-awesome-icon
                  :icon="['fas', 'times']"
                />
              </b-button>
            </b-col>

            <b-col
              cols="2"
              class="p-0"
            >
              <div
                class=" d-flex align-items-center justify-content-center"
              >
                <b-button
                  v-if="(tab.indexOnMain !== null)"
                  id="popover-edit"
                  size="md"
                  :title="$t('tab.tooltip.edit')"
                  variant="link"
                  @click="editBlock(tab.indexOnMain)"
                >
                  <font-awesome-icon
                    :icon="['far', 'edit']"
                  />
                </b-button>

                <c-input-confirm
                  size="xs"
                  link
                  @confirmed="deleteTab(index)"
                />
              </div>
            </b-col>
          </b-row>
        </transition-group>
      </draggable>
    </div>

    <div
      v-else
      class="text-center pt-5 pb-5"
    >
      <h5>
        {{ $t('tab.noTabs') }}
      </h5>
    </div>

    <b-row class="mt-5">
      <b-col>
        <b-button
          v-b-modal.createBlockSelectorTab
          variant="primary"
          :title="$t('tab.tooltip.newBlock')"
        >
          {{ $t('tab.newBlock') }}
        </b-button>
      </b-col>
    </b-row>

    <div class="mt-5 pt-3">
      <h3 class="text-primary">
        {{ $t('tab.displayTitle') }}
      </h3>

      <b-row
        class="mb-3 mt-3 ml-0 mr-0 justify-content-between"
        no-gutters
      >
        <b-form-group label="Appearance">
          <b-form-radio-group
            id="appearance"
            v-model="block.options.tabStyle.appearance"
            buttons
            button-variant="outline-primary"
            size="sm"
            :options="tabStyle.appearance"
            name="display-options"
          />
        </b-form-group>

        <b-form-group label="Alignment">
          <b-form-radio-group
            id="alignment"
            v-model="block.options.tabStyle.alignment"
            buttons
            button-variant="outline-primary"
            size="sm"
            :options="tabStyle.alignment"
            name="display-options"
          />
        </b-form-group>
        <b-form-group label="Fill or Justify">
          <b-form-radio-group
            id="fillJustify"
            v-model="block.options.tabStyle.fillJustify"
            buttons
            button-variant="outline-primary"
            size="sm"
            :options="tabStyle.fillJustify"
            name="display-options"
          />
        </b-form-group>

        <b-form-group label="Vertical or Horizontal">
          <b-form-radio-group
            id="verticalHorizontal"
            v-model="block.options.tabStyle.verticalHorizontal"
            buttons
            button-variant="outline-primary"
            size="sm"
            :options="tabStyle.verticalHorizontal"
            name="display-options"
          />
        </b-form-group>

        <b-form-group label="Tab Position">
          <b-form-radio-group
            id="tabPosition"
            v-model="block.options.tabStyle.tabPosition"
            buttons
            button-variant="outline-primary"
            size="sm"
            :options="tabStyle.tabPosition"
            name="display-options"
          />
        </b-form-group>
      </b-row>
    </div>
    <div class="w-501 divider" />
    <template>
      <div
        class="preview bg-white position-absolute p-3"
      >
        <!-- locale here -->
        <h6 class="text-primary">
          {{ $t('tab.preview') }}
        </h6>

        <b-tabs
          v-bind="{
            align: block.options.tabStyle.alignment,
            fill: block.options.tabStyle.fillJustify === 'fill',
            justified: block.options.tabStyle.fillJustify === 'justified',
            pills: block.options.tabStyle.appearance === 'pills',
            tabs: block.options.tabStyle.appearance === 'tabs',
            small: block.options.tabStyle.appearance === 'small',
            vertical: block.options.tabStyle.verticalHorizontal === 'vertical',
            end: block.options.tabStyle.tabPosition === 'end',
          }"
        >
          <b-tab
            v-for="(tab, index) in block.options.tabs"
            :key="index"
            :title="tab.block.title"
            active
          />
        </b-tabs>
      </div>
    </template>

    <b-modal
      id="createBlockSelectorTab"
      size="lg"
      scrollable
      hide-footer
      :title="$t('tab.newBlockModal')"
    >
      <new-block-selector
        :record-page="!!module"
        @select="addBlock"
      />
    </b-modal>
  </b-tab>
</template>
<script>
import base from './base'
import draggable from 'vuedraggable'

export default {
  i18nOptions: {
    namespaces: 'block',
  },

  name: 'TabConfigurator',
  components: {
    //  Importing like this because configurator is recursive
    NewBlockSelector: () => import('corteza-webapp-compose/src/components/Admin/Page/Builder/Selector'),
    draggable,
  },

  extends: base,

  data () {
    return {
      selectedBlocks: {},
      alert: false,
      tabWarning: false,
      msg: '',
      editFocused: false,
      tabStyle: {
        appearance: [
          { text: this.$t('tab.tabStyle.appearance.tabs'), value: 'tabs' },
          { text: this.$t('tab.tabStyle.appearance.pills'), value: 'pills' },
          { text: this.$t('tab.tabStyle.appearance.small'), value: 'small' },
        ],

        alignment: [
          { text: this.$t('tab.tabStyle.alignment.left'), value: 'left' },
          { text: this.$t('tab.tabStyle.alignment.center'), value: 'center' },
          { text: this.$t('tab.tabStyle.alignment.right'), value: 'right' },
        ],

        fillJustify: [
          { text: this.$t('tab.tabStyle.fillJustify.fill'), value: 'fill' },
          { text: this.$t('tab.tabStyle.fillJustify.justified'), value: 'justified' },
          { text: this.$t('tab.tabStyle.fillJustify.none'), value: '' },
        ],

        verticalHorizontal: [
          { text: this.$t('tab.tabStyle.verticalHorizontal.vertical'), value: 'vertical' },
          { text: this.$t('tab.tabStyle.verticalHorizontal.horizontal'), value: 'none' },
        ],

        tabPosition: [
          { text: this.$t('tab.tabStyle.tabPosition.top'), value: '' },
          { text: this.$t('tab.tabStyle.tabPosition.bottom'), value: 'end' },
        ],
      },
      tabMode: [],
    }
  },

  computed: {

    options () {
      return this.page.blocks.flatMap((block, index) => {
        if (block.kind === 'Tab') {
          return []
        }

        return { value: index, text: block.title || `Block-${block.kind}-${index}` }
      })
    },
  },

  created () {
    this.$root.$on('tab-newBlockMade', this.handleNewBlock)
  },

  destroyed () {
    this.$root.$off('tab-newBlockMade', this.handleNewBlock)
  },

  mounted () {
    this.configureModes()
  },

  methods: {

    Add () {
      this.block.options.tabs.push({
        block: {},
        indexOnMain: null,
      })
      // dragging while adding a new tab causes ui distortions
      this.editFocused = true
      // controls whether save and close button should be active
      this.$root.$emit('tab-checkState')
    },

    createTab (tabIndex) {
      const tab = this.block.options.tabs[tabIndex]
      const blockToTab = this.selectedBlocks[tabIndex]

      if (!this.page.blocks[blockToTab].title) {
        this.tabWarning = true
        this.msg = this.$t('tab.alertTitle')
        setTimeout(() => {
          this.tabWarning = false
        }, 3000)
        return
      }

      if (tab.indexOnMain !== null) {
        // attempt to remove the tabbed block from page but first
        // we check if it is only here it is tabbed before freeing it
        this.untabBlock(tab)
      }

      const newTab = {
        block: this.page.blocks[blockToTab],
        indexOnMain: blockToTab,
      }
      this.updateTabs(newTab, tabIndex)
      this.$root.$emit('tab-checkState')
      this.editFocused = false
    },

    updateTabs (tab, tabIndex) {
      if (tab.block.kind === 'Tab') {
        return
      }
      this.block.options.tabs[tabIndex] = tab
      this.page.blocks[tab.indexOnMain].options.tabbed = true
      this.switchTabMode(tabIndex, 'view')
    },

    deleteTab (tabIndex) {
      const tab = this.block.options.tabs[tabIndex]
      if (tab.indexOnMain !== null) {
        this.untabBlock(tab)
      }
      this.tabMode.splice(tabIndex, 1)
      this.$delete(this.selectedBlocks, tabIndex)
      this.configureModes()
      this.block.options.tabs.splice(tabIndex, 1)
      this.$root.$emit('tab-checkState')
      this.editFocused = false
    },

    untabBlock (tab) {
      const tabOccurrence = this.determineTabOccurrence(tab)
      if (tabOccurrence === 1) {
        this.page.blocks[tab.indexOnMain].options.tabbed = false
      }
    },

    editBlock (index = undefined) {
      const isTabAddedToPage = this.page.blocks.filter(({ kind }) => kind === 'Tab')
        .find(({ options }) => options.blockIndex === this.block.options.blockIndex)
      if (!isTabAddedToPage) {
        this.alert = true
        this.msg = this.$t('tab.alertSave')
        setTimeout(() => {
          this.alert = false
        }, 1500)
        return
      }
      this.$root.$emit('tab-editRequest', index)
    },

    configureModes () {
      this.block.options.tabs.map((tab, index) => {
        if (tab.block.title && tab.indexOnMain !== null) {
          this.$set(this.tabMode, index, 'view')
        }
      })
    },

    handleNewBlock (block) {
      this.Add()
      const getLatestTab = this.block.options.tabs.length - 1
      this.selectedBlocks[getLatestTab] = block
      const blockIndex = this.page.blocks.indexOf(block)
      this.page.blocks[block].title = `Block-${this.page.blocks[block].kind}${blockIndex}`
      this.createTab(getLatestTab)
    },

    switchTabMode (index, mode) {
      this.$set(this.tabMode, index, mode)
    },

    addBlock (block, index = undefined) {
      this.$bvModal.hide('createBlockSelectorTab')
      this.$root.$emit('tab-newBlockRequest', { block, index })
    },

    determineTabOccurrence (tab) {
      const allTabs = this.page.blocks.filter(({ kind, options }) => kind === 'Tab' && options.blockIndex !== this.block.options.blockIndex)
        .concat(this.block)
        .map(({ options }) => options.tabs).flat().map(({ indexOnMain }) => indexOnMain)
      const tabOccurrence = allTabs.filter((t) => t === tab.indexOnMain).length
      return tabOccurrence
    },
  },
}
</script>
<style lang="scss" scoped>

.custom-warning {
  color: red !important;
}

.preview {
  bottom: 0;
  left: 0;
  z-index: 2;
  width: 100%;
  box-shadow: 0 -0.25rem 1rem rgb(0 0 0 / 15%);
}

.divider{
  margin-bottom: 15rem;
}
</style>
