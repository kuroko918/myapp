import firebase from '../plugins/firebase'
import { IUser } from '../types/user'

const uploadImage = async (image: string, userId: IUser['id']): Promise<string> => {
  const storage = firebase.storage()
  const ref = storage.ref().child(`images/users/${userId}/thumbnail`)
  const imageSnap = await ref.putString(image.split(',')[1], 'base64')
  const imageURL = await imageSnap.ref.getDownloadURL()
  return imageURL
}

export default uploadImage
