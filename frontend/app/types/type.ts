export type postType = {
  id: string;
  title: string;
  userId: string;
  category: string;
  tags: string[];
  createdAt: Date;
};

export type userType = {
  id?: string;
  name: string;
  email: string;
  password?: string;
  majors: {
    name: string;
    faculty: string;
  };
  imageUrl: string;
};

export type bookmarkType = {
  id: string;
  userId: string;
  postId: string;
};
