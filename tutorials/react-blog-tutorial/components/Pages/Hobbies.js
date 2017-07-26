import React, { Component } from 'react'
import _ from 'lodash'
import config from '../../config'

import Header from '../Partials/Header'
import HobbyList from '../Partials/HobbyList'
import HobbySingle from '../Partials/HobbySingle'

import AppDispatcher from '../../dispatcher/AppDispatcher'

export default class Hobbies extends Component {

  componentWillMount(){
    this.getPageData()
  }

  componentDidMount(){
    const data = this.props.data
    document.title = config.site.title + ' | ' + data.page.title
  }

  getPageData(){
    AppDispatcher.dispatch({
      action: 'get-page-data',
      page_slug: 'hobby',
      post_slug: this.props.params.slug
    })
  }

  getMoreHobbyItems(){
    AppDispatcher.dispatch({
      action: 'get-more-items'
    })
  }

  render(){

    const data = this.props.data
    const globals = data.globals
    const pages = data.pages
    let main_content

    if(!this.props.params.slug){

      main_content = <HobbyList getMoreHobbyItems={ this.getMoreHobbyItems } data={ data }/>

    } else {

      const hobby_items = data.hobby_items

      const slug = this.props.params.slug
      const hobby_items_object = _.indexBy(hobby_items, 'slug')
      const hobby_item = hobby_items_object[slug]

      main_content = <HobbySingle data={ data } hobby_item={ hobby_item }/>

    }

    return (
      <div>
        <Header data={ data }/>
          <div id="main-content" className="container">
            <div className="row">
              <div className="col-lg-8 col-lg-offset-2 col-md-10 col-md-offset-1">
              { main_content }
              </div>
          </div>
        </div>
      </div>
    )
  }
}
