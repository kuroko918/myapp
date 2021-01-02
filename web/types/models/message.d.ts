import { IUser } from './user'

export type IMessage = {
  id: number
  content: string
  userId: string
  createdAt: string
  updatedAt: string
  deletedAt?: string
  user: IUser
}
