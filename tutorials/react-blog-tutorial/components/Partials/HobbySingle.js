import React, { Component } from 'react'
import _ from 'lodash'
import { Link } from 'react-router'

export default class HobbySingle extends Component {

  render(){

    const hobby_item = this.props.hobby_item

    const style = {
      marginBottom: 20
    }

    return (
      <div>
        <Link to="/hobby" className="btn btn-default" style={ style }>&lt;&lt; Hobby List</Link>
        <h2>{ hobby_item.title }</h2>
        <div dangerouslySetInnerHTML={ {__html: hobby_item.content } }></div>
      </div>
    )
  }
}
