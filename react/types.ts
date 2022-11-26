export interface Message {
  id: number;
  content: string;
  user_id: number;
  channel_id: number;
  time: string;
}

export interface User {
  id: number;
  username: string;
  email: string;
}

export interface MessagesObj {
  messages: Message[];
  status: string;
}
