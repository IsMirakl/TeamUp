import { CanActivate, ExecutionContext, Injectable, UnauthorizedException } from "@nestjs/common";
import { Reflector } from "@nestjs/core";
import { ROLES_KEY } from "./role.decorator";

@Injectable()
export class RoleGuard implements CanActivate {
    constructor(private reflector: Reflector){}

    canActivate(context: ExecutionContext): boolean {
        const requiredRoles = this.reflector.get<string[]>(ROLES_KEY, context.getHandler())

        if(!requiredRoles){
            return true;
        }

        const request = context.switchToHttp().getRequest();
        const user = request.user;

        if(!user){
            throw new UnauthorizedException('User authentication failed');
        }

        if(!user.role){
            throw new UnauthorizedException('User role not found');
        }
        return requiredRoles.includes(user.role);
    }
}