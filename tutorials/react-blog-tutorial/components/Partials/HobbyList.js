import React, { Component } from 'react'
import _ from 'lodash'
import { Link } from 'react-router'

import AppDispatcher from '../../dispatcher/AppDispatcher'

export default class HobbyList extends Component {

  scrollTop(){
    $('html, body').animate({
      scrollTop: $("#main-content").offset().top
    }, 500)
  }

  render(){

    let data = this.props.data
    let item_num = data.item_num
    let hobby_items = data.hobby_items

    let load_more
    let show_more_text = 'Show More Hobbies'

    if(data.loading){
      show_more_text = 'Loading...'
    }

    if(hobby_items && item_num <= hobby_items.length){
      load_more = (
        <div>
          <button className="btn btn-default center-block" onClick={ this.props.getMoreHobbyItems.bind(this) }>
            { show_more_text }
          </button>
        </div>
      )
    }

    hobby_items = _.take(hobby_items, item_num)

    let articles_html = hobby_items.map(( hobby_item ) => {
      let date_obj = new Date(hobby_item.created)
      let created = (date_obj.getMonth()+1) + '/' + date_obj.getDate() + '/' + date_obj.getFullYear()
      return (
        <div key={ 'key-' + hobby_item.slug }>
          <div className="post-preview">
            <h2 className="post-title pointer">
              <Link to={ '/hobbies/' + hobby_item.slug } onClick={ this.scrollTop }>{ hobby_item.title }</Link>
            </h2>
            <p className="post-meta">Posted by <a href="https://cosmicjs.com" target="_blank">Cosmic JS</a> on { created }</p>
          </div>
          <hr/>
        </div>
      )
    })

    return (
      <div>
        <div>{ articles_html }</div>
        { load_more }
      </div>
    )
  }
}
