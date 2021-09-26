class DateSvc {
    public getNowDate() {
        return new Date(new Date().toLocaleString('us-US', { hourCycle: 'h23', timeZone: 'Europe/Paris' }))
    }

    public getNowString() {
        return this.getNowDate().toLocaleString('fr-FR', { hourCycle: 'h23' })
    }
}

const dateSvc = new DateSvc()
export default dateSvc
