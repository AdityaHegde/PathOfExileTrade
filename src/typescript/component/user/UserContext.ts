import {createContext, useEffect, useState} from "react";
import {User} from "@model/User";
import {fetchGet, fetchPost, fetchTry} from "@typescript/utils/fetch";
import {useHistory} from "react-router-dom";

export type UserContextType = {
  user: User;
  login: (userData: any) => Promise<void>;
  signup: (userData: any) => Promise<void>;
  logout: () => Promise<void>;
}

export function useUserContext(): UserContextType {
  const history = useHistory();

  useEffect(() => {
    fetchTry(fetchGet("/auth/user"))
      .then(authRespJson => {
        setUserContextTypeWrapper(new User(authRespJson.data));
        history.push("/dash");
      });
  }, []);

  const logout = async () => {
    await fetchTry(fetchGet("/auth/logout"));
    setUserContextTypeWrapper(null);
  };

  const authenticate = async (userData: any, endpoint: string) => {
    const authRespJson = await fetchTry(fetchPost(`/auth/${endpoint}`, {
      data: {
        id: userData.name,
        type: "users",
        attr: {
          ...userData
        },
      }
    }));
    setUserContextTypeWrapper(new User(authRespJson.data));
    history.push("/dash");
  };

  const login = (userData: any) => {
    return authenticate(userData, "login");
  };

  const signup = (userData: any) => {
    return authenticate(userData, "signup");
  };

  const [userContextType, setUserContextType] = useState<UserContextType>({
    user: null,
    login,
    signup,
    logout,
  });

  const setUserContextTypeWrapper = (user: User) => {
    setUserContextType({
      user,
      login,
      signup,
      logout,
    });
  }

  return userContextType;
}

export const UserContext = createContext<UserContextType>(null);
