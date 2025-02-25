import { CoreApi, SessionUser } from "@goauthentik/api";
import { i18n } from "@lingui/core";
import { DEFAULT_CONFIG } from "./Config";

let globalMePromise: Promise<SessionUser>;
export function me(): Promise<SessionUser> {
    if (!globalMePromise) {
        globalMePromise = new CoreApi(DEFAULT_CONFIG).coreUsersMeRetrieve().then((user) => {
            if (!user.user.settings || !("locale" in user.user.settings)) {
                return user;
            }
            const locale = user.user.settings.locale;
            if (locale && locale !== "") {
                console.debug(`authentik/locale: Activating user's configured locale '${locale}'`);
                i18n.activate(locale);
            }
            return user;
        }).catch((ex) => {
            const defaultUser: SessionUser = {
                user: {
                    pk: -1,
                    isSuperuser: false,
                    isActive: true,
                    groups: [],
                    avatar: "",
                    uid: "",
                    username: "",
                    name: ""
                }
            };
            if (ex.status === 401 || ex.status === 403) {
                window.location.assign("/");
            }
            return defaultUser;
        });
    }
    return globalMePromise;
}
