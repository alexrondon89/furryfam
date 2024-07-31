import jenkins.model.*
import hudson.security.*
import org.jenkinsci.plugins.plaincredentials.impl.StringCredentialsImpl
import com.cloudbees.plugins.credentials.*
import com.cloudbees.plugins.credentials.domains.*

def instance = Jenkins.getInstance()

// Deshabilitar el asistente de instalación
instance.setInstallState(InstallState.INITIAL_SETUP_COMPLETED)

// Configurar seguridad básica
def hudsonRealm = new HudsonPrivateSecurityRealm(false)
def user = hudsonRealm.createAccount("admin", "adminPassword")
instance.setSecurityRealm(hudsonRealm)

def strategy = new FullControlOnceLoggedInAuthorizationStrategy()
instance.setAuthorizationStrategy(strategy)

// Opcional: Configurar un token de API para el usuario
CredentialsStore store = Jenkins.instance.getExtensionList('com.cloudbees.plugins.credentials.SystemCredentialsProvider')[0].getStore()
def token = new StringCredentialsImpl(CredentialsScope.GLOBAL, "admin-token", "description", Secret.fromString("sometoken"))
store.addCredentials(Domain.global(), token)

instance.save()
