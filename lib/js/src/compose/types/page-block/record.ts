import { PageBlock, PageBlockInput, Registry } from './base'

const kind = 'Record'

interface Options {
  fields: unknown[];
  tabbed: boolean;
}

const defaults: Readonly<Options> = Object.freeze({
  fields: [],
  tabbed: false
})

export class PageBlockRecord extends PageBlock {
  readonly kind = kind

  options: Options = { ...defaults }

  constructor (i?: PageBlockInput) {
    super(i)
    this.applyOptions(i?.options as Partial<Options>)
  }

  applyOptions (o?: Partial<Options>): void {
    if (!o) return

    if (o.fields) {
      this.options.fields = o.fields
    }

    if (o.tabbed) {
      this.options.tabbed = o.tabbed
    }
  }
}

Registry.set(kind, PageBlockRecord)
