<template>
  <b-tab
    :title="$t('record.label')"
  >
    <fieldset
      class="form-group"
    >
      <label>
        {{ $t('general.module') }}
      </label>
      <input
        v-if="module"
        v-model="module.name"
        class="form-control"
        type="text"
        readonly
      >
    </fieldset>

    <b-form-group
      :label="$t('module:general.fields')"
    >
      <field-picker
        v-if="module"
        :module="module"
        :fields.sync="options.fields"
        style="max-height: 52vh;"
      />
    </b-form-group>

    <div>
      <b-form-group
        :label="$t('sanitizers.label')"
        label-size="lg"
      >
        <template #label>
          <div
            class="d-flex"
          >
            Condition Fields

            <b-button
              variant="link"
              class="p-0 ml-1 mr-auto"
              @click="add"
            >
              Add +
            </b-button>

            <b-button
              variant="link"
              :href="`${documentationURL}#value-sanitizers`"
              target="_blank"
              class="p-0 ml-1"
            >
              Examples
            </b-button>
          </div>
        </template>
      </b-form-group>

      <b-form-row
        v-for="(fc, i) in options.fieldConditions"
        :key="i"
        :label="`Condition ${i + 1}`"
      >
        <b-input-group
          class="mb-2"
          size="md"
        >
          <vue-select
            v-model="fc.field"
            :get-option-label="f => f.label"
            :placeholder="''"
            :options="fieldOptions"
            :reduce="f => f"
            class="bg-white"
            @close="(() => $root.$emit('checkState', { items: options.fieldConditions, indicator: 'field' }))"
          />
          <b-form-input
            v-model="fc.condition"
            placeholder="When should this field be shown?"
          />
          <c-input-confirm
            class="mt-2 ml-2"
            @confirmed="deleteCondition(i)"
          />
        </b-input-group>
      </b-form-row>
    </div>
  </b-tab>
</template>
<script>
import base from './base'
import FieldPicker from 'corteza-webapp-compose/src/components/Common/FieldPicker'
import { VueSelect } from 'vue-select'

export default {
  i18nOptions: {
    namespaces: 'block',
  },

  name: 'Record',

  components: {
    FieldPicker,
    VueSelect,
  },

  extends: base,

  computed: {
    documentationURL () {
      // eslint-disable-next-line no-undef
      const [year, month] = VERSION.split('.')
      return `https://docs.cortezaproject.org/corteza-docs/${year}.${month}/integrator-guide/compose-configuration/index.html`
    },

    fieldOptions () {
      const nonRequiredFields = this.module.fields.filter(f => !f.isRequired)
      const unConditionedFields = nonRequiredFields.filter(f => !this.options.fieldConditions.find(fc => fc.field.name === f.name))
      const fields = unConditionedFields.map(f => {
        return {
          name: f.name,
          label: f.label,
          fieldID: f.fieldID,
          isRequired: f.isRequired,
        }
      })
      return fields
    },
  },

  methods: {
    add () {
      if (this.options.fieldConditions.length >= this.module.fields.length) return
      this.options.fieldConditions.push({
        field: {},
        condition: '',
      })
      this.$root.$emit('checkState', { items: this.options.fieldConditions, indicator: 'field' })
    },

    deleteCondition (i) {
      this.options.fieldConditions.splice(i, 1)
      this.$root.$emit('checkState', { items: this.options.fieldConditions, indicator: 'field' })
    },
  },
}
</script>

<style lang="scss" scoped>
  div.fields {
    display: flex;
    flex-flow: row nowrap;

    & > div {
      flex: 1;
      margin: 5px;

      button.all {
        float: right;
        font-size: 80%;
      }

      .drag-area {
        height: 150px;
        overflow-x: auto;
        border: 1px solid silver;
        padding: 2px;
      }
    }
  }
</style>
