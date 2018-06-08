import pytest
import subprocess
import testinfra


# scope='session' uses the same container for all the tests;
# scope='function' uses a new container per test function.
@pytest.fixture(scope='session')
def host(request):
    # build local ./Dockerfile
    # subprocess.check_call(['docker', 'build', '-t', 'myimage', '.'])
    # run a container
    docker_id = subprocess.check_output(
        ['docker', 'run', '-d', 'local/snippet-new-platform', 'sleep', '1d']).decode().strip()
    subprocess.check_output(['docker', 'cp', '.', docker_id+':/tmp/src'])
    
    # return a testinfra connection to the container
    yield testinfra.get_host("docker://" + docker_id)
    # at the end of the test suite, destroy the container
    subprocess.check_call(['docker', 'rm', '-f', docker_id])


def test_myimage(host):
    # 'host' now binds to the container
    host.run_expect([0], '/usr/libexec/s2i/assemble')
    assert host.package("python3").is_installed
    assert host.check_output('echo \'oi\'') == 'oi'

def test_passwd_file(host):
    passwd = host.file("/etc/passwd")
    assert passwd.contains("root")
    assert passwd.user == "root"
    assert passwd.group == "root"
    assert passwd.mode == 0o644


# def test_nginx_is_installed(host):
#     nginx = host.package("nginx")
#     assert nginx.is_installed
#     assert nginx.version.startswith("1.2")


# def test_nginx_running_and_enabled(host):
#     nginx = host.service("nginx")
#     assert nginx.is_running
#     assert nginx.is_enabled