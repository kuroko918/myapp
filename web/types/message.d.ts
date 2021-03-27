import { IUser } from './user'

export type IMessage = {
  id: string
  content: string
  userId: string
  createdAt: string
  updatedAt: string
  user: IUser
}
