import {
  getAuth,
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
  // signOut,
  sendPasswordResetEmail,
  verifyBeforeUpdateEmail,
  updatePassword,
  signInWithPopup,
} from "firebase/auth";
import app from "./firebase";
import { GoogleAuthProvider } from "firebase/auth/web-extension";

export interface FireBaseAuth {
  signUp: (email: string, password: string) => Promise<any>;
  signIn: (email: string, password: string) => Promise<any>;
  resetPassword: (email: string) => Promise<any>;
  updateUserEmail: (email: string, userInfo: any) => Promise<any>;
  updateUserPassword: (password: string, userInfo: any) => Promise<any>;
  signInWithGoogle: () => Promise<any>;
}

export const auth = getAuth(app);
const signUp = (email: string, password: string) => {
  return createUserWithEmailAndPassword(auth, email, password);
};

const signIn = (email: string, password: string) => {
  return signInWithEmailAndPassword(auth, email, password);
};

const resetPassword = (email: string) => {
  return sendPasswordResetEmail(auth, email);
};

const updateUserEmail = (email: string, userInfo: any) => {
  return verifyBeforeUpdateEmail(userInfo, email);
};

const updateUserPassword = (password: string, userInfo: any) => {
  return updatePassword(userInfo, password);
};

const signInWithGoogle = () => {
  return signInWithPopup(auth, new GoogleAuthProvider());
};

const value = {
  signUp,
  signIn,
  resetPassword,
  updateUserEmail,
  updateUserPassword,
  signInWithGoogle,
};

export default value;
