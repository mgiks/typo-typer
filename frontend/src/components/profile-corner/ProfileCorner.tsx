import { faker } from '@faker-js/faker'
import { useAppDispatch } from '../../hooks'
import { setName } from '../../slices/multiplayerData.slice'

function ProfileCorner() {
  const name = faker.word.adjective() + ' ' + faker.word.noun()
  useAppDispatch()(setName(name))
  return <button>{name}</button>
}

export default ProfileCorner
