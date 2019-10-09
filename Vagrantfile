Vagrant.configure("2") do |config|
  config.vm.provider "virtualbox" do |v|
    v.memory = "1024"
    v.cpus = 2
    v.customize ["modifyvm", :id, "--cpuexecutioncap", "70"]
  end

  config.vm.define "gomongoRH7" do |gomongoRH7|
#   gomongoRH7.vm.box = "generic/rhel7"
    gomongoRH7.vm.box = "clouddood/RH7.5_baserepo"
    #gomongoRH7.vm.box = "javier-lopez/rhel-7.4"
    #gomongoRH7.vm.box = "xianlin/rhel-7.4"
    gomongoRH7.vm.hostname = "gomongoRH7"
    gomongoRH7.vm.network "private_network", ip: "192.168.60.182"
    gomongoRH7.vm.provision "shell", :inline => "sudo echo '192.168.60.182 gomongoRH7.local gomongoRH7' >> /etc/hosts"
         
    gomongoRH7.vm.provision "ansible" do |ansible|
      ansible.playbook = "deploy_gomongoRH7.yml"
      ansible.inventory_path = "vagrant_hosts"
      #ansible.tags = ansible_tags
      #ansible.verbose = ansible_verbosity
      #ansible.extra_vars = ansible_extra_vars
      #ansible.limit = ansible_limit
    end
  end
  config.vm.define "gomysqlRH7" do |gomongoRH7|
#   gomysqlRH7.vm.box = "generic/rhel7"
#   gomysqlRH7.vm.box = "iamseth/rhel-7.3"
    gomysqlRH7.vm.box = "clouddood/RH7.5_baserepo"
    #gomysqlRH7.vm.box = "javier-lopez/rhel-7.4"
    #gomysqlRH7.vm.box = "xianlin/rhel-7.4"
    gomysqlRH7.vm.hostname = "gomysqlRH7"
    gomysqlRH7.vm.network "private_network", ip: "192.168.60.183"
    gomysqlRH7.vm.provision "shell", :inline => "sudo echo '192.168.60.183 gomysqlRH7.local gomysqlRH7' >> /etc/hosts"
         
    gomysqlRH7.vm.provision "ansible" do |ansible|
      ansible.playbook = "deploy_gomysqlRH7.yml"
      ansible.inventory_path = "vagrant_hosts"
      #ansible.tags = ansible_tags
      #ansible.verbose = ansible_verbosity
      #ansible.extra_vars = ansible_extra_vars
      #ansible.limit = ansible_limit
    end
  end
end
