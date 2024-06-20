import * as React from "react";
import "firebase/compat/auth";
import { FireBaseAuth } from "../firebase/firebaseUtil";
import firebaseUtil from "../firebase/firebaseUtil";
import { auth } from "../firebase/firebaseUtil";

const AuthContext = React.createContext<AuthContextType | undefined>(undefined);

export const useAuth = () => {
  const context = React.useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

interface AuthContextType {
  userInfo: any;
  isAuthenticated: boolean;
  firebaseUtil: FireBaseAuth;
}
const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [userInfo, setUserInfo] = React.useState<any>(null);
  const [isAuthenticated, setIsAuthenticated] = React.useState<boolean>(false);
  React.useEffect(() => {
    const unsubscribe = auth.onAuthStateChanged(async (user) => {
      if (user) {
        //call BE for login
        setUserInfo(user);
        setIsAuthenticated(true);
      } else {
        setUserInfo(null);
        setIsAuthenticated(false);
      }
    });
    return () => unsubscribe();
  }, []);
  return (
    <AuthContext.Provider
      value={{
        userInfo,
        isAuthenticated,
        firebaseUtil: firebaseUtil as FireBaseAuth,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
