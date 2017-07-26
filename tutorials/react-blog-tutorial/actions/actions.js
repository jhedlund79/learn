import config from '../config'
import Cosmic from 'cosmicjs'
import _ from 'lodash'

import AppStore from '../stores/AppStore'

export function getStore(callback){

  let pages = {}

  Cosmic.getObjects(config, function(err, response){

    let objects = response.objects

    let globals = AppStore.data.globals
    globals.text = response.object['text']
    let metafields = globals.text.metafields
    let menu_title = _.findWhere(metafields, { key: 'menu-title'})
    globals.text.menu_title = menu_title.value

    let footer_text = _.findWhere(metafields, { key: 'footer-text' })
    globals.text.footer_text = "Footer stuff"

    let site_title = _.findWhere(metafields, { key: 'site-title' })
    globals.text.site_title= site_title.value

    globals.social = response.object['social']
    metafields = globals.social.metafields
    let twitter = _.findWhere(metafields, { key: 'twitter' })
    globals.social.twitter = twitter.value
    let facebook = _.findWhere(metafields, { key: 'facebook' })
    globals.social.facebook = facebook.value
    let github = _.findWhere(metafields, { key: 'github' })
    globals.social.github = github.value

    const nav_items = response.object['nav'].metafields
    globals.nav_items = nav_items

    AppStore.data.globals = globals

    let pages = objects.type.page
    AppStore.data.pages = pages

    let articles = objects.type['post']
    articles = _.sortBy(articles, 'order')
    AppStore.data.articles = articles

    let hobby_items = objects.type['hobby']
    hobby_items = _.sortBy(hobby_items, 'order')
    AppStore.data.hobby_items = hobby_items

    AppStore.data.ready = true
    AppStore.emitChange()

    if(callback){
      callback(false, AppStore)
    }

  })
}

export function getPageData(page_slug, post_slug){

   if(!page_slug || page_slug === 'blog')
     page_slug = 'home'

   const data = AppStore.data
   const pages = data.pages
   const page = _.findWhere(pages, { slug: page_slug })
   const metafields = page.metafields
   if(metafields){
     const hero = _.findWhere(metafields, { key: 'hero' })
     page.hero = config.bucket.media_url + '/' + hero.value

     const headline = _.findWhere(metafields, { key: 'headline' })
     page.headline = headline.value

     const subheadline = _.findWhere(metafields, { key: 'subheadline' })
     page.subheadline = subheadline.value
   }

   if(post_slug){
     if(page_slug === 'home'){
       const articles = data.articles
       const article = _.findWhere(articles, { slug: post_slug })
       page.title = article.title
     }
     if(page_slug === 'hobby'){
       const hobby_items = data.hobby_items
       const hobby_item = _.findWhere(hobby_items, { slug: post_slug })
       page.title = hobby_item.title
     }
   }
   AppStore.data.page = page
   AppStore.emitChange()
}

export function getMoreItems(){

   AppStore.data.loading = true
   AppStore.emitChange()

   setTimeout(function(){
     let item_num = AppStore.data.item_num
     let more_item_num = item_num + 5
     AppStore.data.item_num = more_item_num
     AppStore.data.loading = false
     AppStore.emitChange()
   }, 300)
}
