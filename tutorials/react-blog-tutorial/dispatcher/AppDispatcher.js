import { Dispatcher } from 'flux'
import { getStore, getPageData, getMoreItems } from '../actions/actions'

const AppDispatcher = new Dispatcher()

AppDispatcher.register((payload) => {
  let action = payload.action

  switch(action){
    case 'get-app-store':
      getStore()
      break
    case 'get-page-data':
      getPageData()
      break
    case 'get-more-items':
      getPageData()
      break
    default:
      return true

  }

  return true

})

export default AppDispatcher
