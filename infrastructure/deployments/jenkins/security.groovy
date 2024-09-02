import jenkins.model.*
import hudson.security.*
import hudson.model.*

def instance = Jenkins.get()

// Configurar el sistema de seguridad
def hudsonRealm = new HudsonPrivateSecurityRealm(false)
hudsonRealm.createAccount('admin', 'admin')
instance.setSecurityRealm(hudsonRealm)

def strategy = new FullControlOnceLoggedInAuthorizationStrategy()
strategy.setAllowAnonymousRead(false)
instance.setAuthorizationStrategy(strategy)

// Guardar la configuración
instance.save()
