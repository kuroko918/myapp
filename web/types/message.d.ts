import { IUser } from './user'

export interface IMessage {
  id: string;
  content: string;
  userId: string;
  createdAt: string;
  updatedAt: string;
  user: IUser;
}
